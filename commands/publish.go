package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/usk81/easyindex"

	"github.com/usk81/easyindex-cli/errors"
	"github.com/usk81/easyindex-cli/usecase"
)

var (
	publishCmd = &cobra.Command{
		Use:   "publish",
		Short: "Notifies that a URL has been updated or deleted.",
		Long:  "Notifies that a URL has been updated or deleted.",
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

func publishCommandWithUsecase(cmd *cobra.Command, args []string, nt easyindex.NotificationType, f usecase.PublishFunc) error {
	limit, err := GetIntFlagOrEnv(cmd.Flags(), flagKeyLimit, envKeyLimit, defaultLimit)
	if err != nil {
		return errors.NewArgError(err, "int")
	}
	skip, err := GetBoolFlagOrEnv(cmd.Flags(), flagKeySkip, envKeySkip, defaultSkip)
	if err != nil {
		return errors.NewArgError(err, "bool")
	}
	cred, err := GetStringFlagOrEnv(cmd.Flags(), flagKeyCredentials, envKeyCredentials, defaultCredentials)
	if err != nil {
		return errors.NewArgError(err, "string")
	}
	ignore, err := GetBoolFlagOrEnv(cmd.Flags(), flagKeyIgnore, envKeyIgnore, defaultIgnore)
	if err != nil {
		return errors.NewArgError(err, "bool")
	}

	result, err := f(nt, args, cred, limit, skip, ignore)
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
	return publishCommandWithUsecase(cmd, args, easyindex.NotificationTypeUpdated, usecase.Publish)
}

func publishDeletedCommand(cmd *cobra.Command, args []string) error {
	return publishCommandWithUsecase(cmd, args, easyindex.NotificationTypeDeleted, usecase.Publish)
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
	RootCmd.AddCommand(publishCmd)
}
