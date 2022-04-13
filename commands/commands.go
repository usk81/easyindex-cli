package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	// RootCmd sets task command config
	RootCmd = &cobra.Command{
		Use: "easyindex",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage() // nolint
		},
	}
)

// Run runs CLI action
func Run() {
	if err := RootCmd.Execute(); err != nil {
		Exit(err, 1)
	}
}

// Exit finishs requests
func Exit(err error, codes ...int) {
	var code int
	if len(codes) > 0 {
		code = codes[0]
	} else {
		code = 2
	}
	fmt.Println(err)
	os.Exit(code)
}

// GetBoolFlagOrEnv ...
func GetBoolFlagOrEnv(fs *pflag.FlagSet, name string, envKey string, def bool) (bool, error) {
	v, err := fs.GetBool(name)
	if err != nil {
		return false, err
	}
	if v == def {
		if ev := os.Getenv(envKey); ev != "" {
			b, err := strconv.ParseBool(ev)
			if err != nil {
				return false, err
			}
			return b, nil
		}
	}
	return v, nil
}

// GetIntFlagOrEnv ...
func GetIntFlagOrEnv(fs *pflag.FlagSet, name string, envKey string, def int) (int, error) {
	v, err := fs.GetInt(name)
	if err != nil {
		return 0, err
	}
	if v == def {
		if ev := os.Getenv(envKey); ev != "" {
			i, err := strconv.Atoi(ev)
			if err != nil {
				return 0, err
			}
			return i, nil
		}
	}
	return v, nil
}

// GetStringFlagOrEnv ...
func GetStringFlagOrEnv(fs *pflag.FlagSet, name string, envKey string, def string) (string, error) {
	v, err := fs.GetString(name)
	if err != nil {
		return "", err
	}
	if v == def || v == "" {
		if ev := os.Getenv(envKey); ev != "" {
			return ev, nil
		}
	}
	return v, nil
}
