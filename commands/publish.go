package commands

import (
	"encoding/json"
	"fmt"

	"github.com/gocarina/gocsv"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/usk81/aveo"

	"github.com/usk81/easyindex"
	"github.com/usk81/easyindex/coordinator"
	"github.com/usk81/easyindex/logger"

	"github.com/usk81/easyindex-cli/errors"
	"github.com/usk81/easyindex-cli/usecase"
)

var (
	publishCmd = &cobra.Command{
		Use:   "publish",
		Short: "Notifies that a URL has been updated or deleted.",
		Long:  "Notifies that a URL has been updated or deleted.",
		RunE:  publishByFileCommand,
	}

	publishUpdatedCmd = &cobra.Command{
		Use:   "updated",
		Short: "Notifies that a URL has been updated.",
		Long:  "Notifies that a URL has been updated.",
		RunE:  publishUpdatedCommand,
	}

	publishDeletedCmd = &cobra.Command{
		Use:   "deleted",
		Short: "Notifies that a URL has been deleted.",
		Long:  "Notifies that a URL has been deleted.",
		RunE:  publishDeletedCommand,
	}
)

const (
	defaultCredentials = "credentials.json"
	defaultLimit       = 200
	defaultSkip        = false
	defaultIgnore      = false

	envKeyCredentials = "EASYINDEX_CREDENTIALS_PATH"
	envKeyLimit       = "EASYINDEX_REQUEST_LIMIT"
	envKeySkip        = "EASYINDEX_SKIP"
	envKeyIgnore      = "EASYINDEX_IGNORE_PRECHECK"

	flagKeyCredentials = "credentials"
	flagKeyLimit       = "limit"
	flagKeySkip        = "skip"
	flagKeyIgnore      = "ignore"
	flagKeyCSV         = "csv"
	flagKeyJSON        = "json"
)

func printPublishCallResponse(r *usecase.PublishResult) {
	fmt.Println("[result]")
	fmt.Printf("Total: %d\n", r.Total)
	fmt.Printf("Count: %d\n", r.Count)
	if len(r.Skips) > 0 {
		fmt.Println("Skips:")
		for i := range r.Skips {
			v := &r.Skips[i]
			fmt.Printf("  [%s] %s : %s\n", v.NotificationType, v.URL, v.Reason.Error())
		}
	}
}

func publishByFileCommand(cmd *cobra.Command, args []string) error {
	return publishByFileCommandWithFileSystem(cmd, args, aveo.NewOs(), afero.NewOsFs(), usecase.PublishBulk)
}

func publishByFileCommandWithFileSystem(cmd *cobra.Command, args []string, env aveo.Env, fs afero.Fs, f usecase.PublishBulkFunc) error {
	limit, cred, ignore, skip, err := getOptions(env, cmd.Flags())
	if err != nil {
		return err
	}
	rq := []*usecase.PublishRequest{}
	cp, err := cmd.Flags().GetString(flagKeyCSV)
	if err != nil {
		return errors.NewArgError(err, "string")
	}
	if cp != "" {
		bs, err := afero.ReadFile(fs, cp)
		if err != nil {
			return err
		}
		if err = gocsv.UnmarshalBytes(bs, &rq); err != nil {
			return err
		}
	} else {
		jp, err := cmd.Flags().GetString(flagKeyJSON)
		if err != nil {
			return errors.NewArgError(err, "string")
		}
		if jp != "" {
			bs, err := afero.ReadFile(fs, cp)
			if err != nil {
				return err
			}
			if err = json.Unmarshal(bs, &rq); err != nil {
				return err
			}
		}
	}

	l, err := logger.New("debug")
	if err != nil {
		return err
	}
	mgr, err := coordinator.New(coordinator.Config{
		CredentialsFile: &cred,
		IgnorePreCheck:  ignore,
		Skip:            skip,
		Logger:          l,
	})
	if err != nil {
		return err
	}

	result, err := f(mgr, rq, limit)
	if err != nil {
		return err
	}
	if result == nil {
		return fmt.Errorf("unknown error: can not get result")
	}
	printPublishCallResponse(result)
	return nil
}

func publishCommandWithUsecase(cmd *cobra.Command, args []string, env aveo.Env, nt easyindex.NotificationType, f usecase.PublishFunc) error {
	limit, cred, ignore, skip, err := getOptions(env, cmd.Flags())
	if err != nil {
		return err
	}

	l, err := logger.New("debug")
	if err != nil {
		return err
	}
	mgr, err := coordinator.New(coordinator.Config{
		CredentialsFile: &cred,
		IgnorePreCheck:  ignore,
		Skip:            skip,
		Logger:          l,
	})
	if err != nil {
		return err
	}

	result, err := f(mgr, nt, args, limit)
	if err != nil {
		return err
	}
	if result == nil {
		return fmt.Errorf("unknown error: can not get result")
	}
	printPublishCallResponse(result)
	return nil
}

func publishUpdatedCommand(cmd *cobra.Command, args []string) error {
	return publishCommandWithUsecase(cmd, args, aveo.NewOs(), easyindex.NotificationTypeUpdated, usecase.Publish)
}

func publishDeletedCommand(cmd *cobra.Command, args []string) error {
	return publishCommandWithUsecase(cmd, args, aveo.NewOs(), easyindex.NotificationTypeDeleted, usecase.Publish)
}

func getOptions(env aveo.Env, fs *pflag.FlagSet) (limit int, cred string, ignore, skip bool, err error) {
	if limit, err = GetIntFlagOrEnv(env, fs, flagKeyLimit, envKeyLimit, defaultLimit); err != nil {
		return 0, "", false, false, errors.NewArgError(err, "int")
	}
	if skip, err = GetBoolFlagOrEnv(env, fs, flagKeySkip, envKeySkip, defaultSkip); err != nil {
		return 0, "", false, false, errors.NewArgError(err, "bool")
	}
	if cred, err = GetStringFlagOrEnv(env, fs, flagKeyCredentials, envKeyCredentials, defaultCredentials); err != nil {
		return 0, "", false, false, errors.NewArgError(err, "string")
	}
	if ignore, err = GetBoolFlagOrEnv(env, fs, flagKeyIgnore, envKeyIgnore, defaultIgnore); err != nil {
		return 0, "", false, false, errors.NewArgError(err, "bool")
	}
	return
}

func setFlagForPublishCommand(cmd *cobra.Command) {
	flags := cmd.Flags()
	flags.StringP(flagKeyCredentials, "c", defaultCredentials, "Credentials file path")
	flags.IntP(flagKeyLimit, "l", defaultLimit, "Limit the number of API request")
	flags.BoolP(flagKeySkip, "s", defaultSkip, "Skip API request if can not access page")
	flags.BoolP(flagKeyIgnore, "i", defaultIgnore, "Do not pre-check")
}

func init() {
	cmds := []*cobra.Command{
		publishUpdatedCmd,
		publishDeletedCmd,
	}
	for _, cmd := range cmds {
		setFlagForPublishCommand(cmd)
		publishCmd.AddCommand(cmd)
	}
	setFlagForPublishCommand(publishCmd)
	flags := publishCmd.Flags()
	flags.StringP(flagKeyCSV, "C", "", "path of csv file to input")
	flags.StringP(flagKeyJSON, "j", "", "path of json file to input")
	RootCmd.AddCommand(publishCmd)
}
