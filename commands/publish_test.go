package commands

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/usk81/aveo"
)

func Test_getOptions(t *testing.T) {
	type args struct {
		env aveo.Env
		fs  *pflag.FlagSet
	}
	tests := []struct {
		name       string
		args       args
		wantLimit  int
		wantCred   string
		wantIgnore bool
		wantSkip   bool
		wantErr    bool
	}{
		{
			name: "all",
			args: args{
				env: aveo.NewMap(nil),
				fs: (func() *pflag.FlagSet {
					fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
					fs.String(flagKeyCredentials, defaultCredentials, "Credentials file path")
					fs.Int(flagKeyLimit, defaultLimit, "Limit the number of API request")
					fs.Bool(flagKeySkip, defaultSkip, "Skip API request if can not access page")
					fs.Bool(flagKeyIgnore, defaultIgnore, "Do not pre-check")
					fs.Set(flagKeyCredentials, "foobar.json")
					fs.Set(flagKeyLimit, "100")
					fs.Set(flagKeySkip, "true")
					fs.Set(flagKeyIgnore, "true")
					return fs
				})(),
			},
			wantLimit:  100,
			wantCred:   "foobar.json",
			wantIgnore: true,
			wantSkip:   true,
			wantErr:    false,
		},
		{
			name: "none",
			args: args{
				env: aveo.NewMap(nil),
				fs: (func() *pflag.FlagSet {
					fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
					fs.String(flagKeyCredentials, defaultCredentials, "Credentials file path")
					fs.Int(flagKeyLimit, defaultLimit, "Limit the number of API request")
					fs.Bool(flagKeySkip, defaultSkip, "Skip API request if can not access page")
					fs.Bool(flagKeyIgnore, defaultIgnore, "Do not pre-check")
					return fs
				})(),
			},
			wantLimit:  defaultLimit,
			wantCred:   defaultCredentials,
			wantIgnore: defaultIgnore,
			wantSkip:   defaultSkip,
			wantErr:    false,
		},
		{
			name: "credentials",
			args: args{
				env: aveo.NewMap(nil),
				fs: (func() *pflag.FlagSet {
					fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
					fs.String(flagKeyCredentials, defaultCredentials, "Credentials file path")
					fs.Int(flagKeyLimit, defaultLimit, "Limit the number of API request")
					fs.Bool(flagKeySkip, defaultSkip, "Skip API request if can not access page")
					fs.Bool(flagKeyIgnore, defaultIgnore, "Do not pre-check")
					fs.Set(flagKeyCredentials, "foobar.json")
					return fs
				})(),
			},
			wantLimit:  defaultLimit,
			wantCred:   "foobar.json",
			wantIgnore: defaultIgnore,
			wantSkip:   defaultSkip,
			wantErr:    false,
		},
		{
			name: "limit",
			args: args{
				env: aveo.NewMap(nil),
				fs: (func() *pflag.FlagSet {
					fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
					fs.String(flagKeyCredentials, defaultCredentials, "Credentials file path")
					fs.Int(flagKeyLimit, defaultLimit, "Limit the number of API request")
					fs.Bool(flagKeySkip, defaultSkip, "Skip API request if can not access page")
					fs.Bool(flagKeyIgnore, defaultIgnore, "Do not pre-check")
					fs.Set(flagKeyLimit, "100")
					return fs
				})(),
			},
			wantLimit:  100,
			wantCred:   defaultCredentials,
			wantIgnore: defaultIgnore,
			wantSkip:   defaultSkip,
			wantErr:    false,
		},
		{
			name: "skip",
			args: args{
				env: aveo.NewMap(nil),
				fs: (func() *pflag.FlagSet {
					fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
					fs.String(flagKeyCredentials, defaultCredentials, "Credentials file path")
					fs.Int(flagKeyLimit, defaultLimit, "Limit the number of API request")
					fs.Bool(flagKeySkip, defaultSkip, "Skip API request if can not access page")
					fs.Bool(flagKeyIgnore, defaultIgnore, "Do not pre-check")
					fs.Set(flagKeySkip, "true")
					return fs
				})(),
			},
			wantLimit:  defaultLimit,
			wantCred:   defaultCredentials,
			wantIgnore: defaultIgnore,
			wantSkip:   true,
			wantErr:    false,
		},
		{
			name: "ignore",
			args: args{
				env: aveo.NewMap(nil),
				fs: (func() *pflag.FlagSet {
					fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
					fs.String(flagKeyCredentials, defaultCredentials, "Credentials file path")
					fs.Int(flagKeyLimit, defaultLimit, "Limit the number of API request")
					fs.Bool(flagKeySkip, defaultSkip, "Skip API request if can not access page")
					fs.Bool(flagKeyIgnore, defaultIgnore, "Do not pre-check")
					fs.Set(flagKeyIgnore, "true")
					return fs
				})(),
			},
			wantLimit:  defaultLimit,
			wantCred:   defaultCredentials,
			wantIgnore: true,
			wantSkip:   defaultSkip,
			wantErr:    false,
		},
		{
			name: "credentials_empty",
			args: args{
				env: aveo.NewMap(nil),
				fs: (func() *pflag.FlagSet {
					fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
					fs.String(flagKeyCredentials, defaultCredentials, "Credentials file path")
					fs.Int(flagKeyLimit, defaultLimit, "Limit the number of API request")
					fs.Bool(flagKeySkip, defaultSkip, "Skip API request if can not access page")
					fs.Bool(flagKeyIgnore, defaultIgnore, "Do not pre-check")
					fs.Set(flagKeyCredentials, "")
					return fs
				})(),
			},
			wantLimit:  defaultLimit,
			wantCred:   defaultCredentials,
			wantIgnore: defaultIgnore,
			wantSkip:   defaultSkip,
			wantErr:    false,
		},
		{
			name: "env_credentials",
			args: args{
				env: aveo.NewMap(map[string]string{
					envKeyCredentials: "env_credentials.json",
				}),
				fs: (func() *pflag.FlagSet {
					fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
					fs.String(flagKeyCredentials, defaultCredentials, "Credentials file path")
					fs.Int(flagKeyLimit, defaultLimit, "Limit the number of API request")
					fs.Bool(flagKeySkip, defaultSkip, "Skip API request if can not access page")
					fs.Bool(flagKeyIgnore, defaultIgnore, "Do not pre-check")
					return fs
				})(),
			},
			wantLimit:  defaultLimit,
			wantCred:   "env_credentials.json",
			wantIgnore: defaultIgnore,
			wantSkip:   defaultSkip,
			wantErr:    false,
		},
		{
			name: "env_limit",
			args: args{
				env: aveo.NewMap(map[string]string{
					envKeyLimit: "1000",
				}),
				fs: (func() *pflag.FlagSet {
					fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
					fs.String(flagKeyCredentials, defaultCredentials, "Credentials file path")
					fs.Int(flagKeyLimit, defaultLimit, "Limit the number of API request")
					fs.Bool(flagKeySkip, defaultSkip, "Skip API request if can not access page")
					fs.Bool(flagKeyIgnore, defaultIgnore, "Do not pre-check")
					return fs
				})(),
			},
			wantLimit:  1000,
			wantCred:   defaultCredentials,
			wantIgnore: defaultIgnore,
			wantSkip:   defaultSkip,
			wantErr:    false,
		},
		{
			name: "env_ignore",
			args: args{
				env: aveo.NewMap(map[string]string{
					envKeyIgnore: "true",
				}),
				fs: (func() *pflag.FlagSet {
					fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
					fs.String(flagKeyCredentials, defaultCredentials, "Credentials file path")
					fs.Int(flagKeyLimit, defaultLimit, "Limit the number of API request")
					fs.Bool(flagKeySkip, defaultSkip, "Skip API request if can not access page")
					fs.Bool(flagKeyIgnore, defaultIgnore, "Do not pre-check")
					return fs
				})(),
			},
			wantLimit:  defaultLimit,
			wantCred:   defaultCredentials,
			wantIgnore: true,
			wantSkip:   defaultSkip,
			wantErr:    false,
		},
		{
			name: "env_skip",
			args: args{
				env: aveo.NewMap(map[string]string{
					envKeySkip: "true",
				}),
				fs: (func() *pflag.FlagSet {
					fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
					fs.String(flagKeyCredentials, defaultCredentials, "Credentials file path")
					fs.Int(flagKeyLimit, defaultLimit, "Limit the number of API request")
					fs.Bool(flagKeySkip, defaultSkip, "Skip API request if can not access page")
					fs.Bool(flagKeyIgnore, defaultIgnore, "Do not pre-check")
					return fs
				})(),
			},
			wantLimit:  defaultLimit,
			wantCred:   defaultCredentials,
			wantIgnore: defaultIgnore,
			wantSkip:   true,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotLimit, gotCred, gotIgnore, gotSkip, err := getOptions(tt.args.env, tt.args.fs)
			if (err != nil) != tt.wantErr {
				t.Errorf("getOptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotLimit != tt.wantLimit {
				t.Errorf("getOptions() gotLimit = %v, want %v", gotLimit, tt.wantLimit)
			}
			if gotCred != tt.wantCred {
				t.Errorf("getOptions() gotCred = %v, want %v", gotCred, tt.wantCred)
			}
			if gotIgnore != tt.wantIgnore {
				t.Errorf("getOptions() gotIgnore = %v, want %v", gotIgnore, tt.wantIgnore)
			}
			if gotSkip != tt.wantSkip {
				t.Errorf("getOptions() gotSkip = %v, want %v", gotSkip, tt.wantSkip)
			}
		})
	}
}
