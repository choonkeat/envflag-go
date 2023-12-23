package envflag

import (
	"flag"
	"os"
	"strconv"
	"strings"
	"time"
)

// Parse parses the command-line flags from os.Args[1:]. Must be called after
// all flags are defined and before flags are accessed by the program.
func Parse() {
	flag.Parse()
}

// Helper function to get a pointer to a value. Bundled here because our
// default values are pointers.
func Ptr[T any](v T) *T { return &v }

// converts flag names like "foo-bar" to env names like "FOO_BAR"
func flagToEnv(name string) string {
	return strings.ToUpper(strings.Replace(name, "-", "_", -1))
}

func flagUsage(name string, usage string) string {
	return usage + " (env " + flagToEnv(name) + ")"
}

// Just like flag.BoolVar, but allows the default value to be set from an
// environment variable named in the format of FLAG_NAME.
//
// This function panics if no default value is provided and the environment
// variable is invalid.
func BoolVar(p *bool, name string, value *bool, usage string) {
	switch os.Getenv(flagToEnv(name)) {
	case "1", "t", "T", "true", "TRUE", "True":
		flag.BoolVar(p, name, true, flagUsage(name, usage))
		return
	case "0", "f", "F", "false", "FALSE", "False":
		flag.BoolVar(p, name, false, flagUsage(name, usage))
		return
	case "":
		if value != nil {
			flag.BoolVar(p, name, *value, flagUsage(name, usage))
			return
		}
	}
	panic("invalid default boolean value for env " + flagToEnv(name))
}

// Just like flag.DurationVar, but allows the default value to be set from an
// environment variable named in the format of FLAG_NAME.
//
// This function panics if no default value is provided and the environment
// variable is invalid.
func DurationVar(p *time.Duration, name string, value *time.Duration, usage string) {
	switch s := os.Getenv(flagToEnv(name)); s {
	case "":
		if value != nil {
			flag.DurationVar(p, name, *value, flagUsage(name, usage))
			return
		}
	default:
		if d, err := time.ParseDuration(s); err == nil {
			flag.DurationVar(p, name, d, flagUsage(name, usage))
			return
		}
	}
	panic("invalid default duration value for env " + flagToEnv(name))
}

// Just like flag.Float64Var, but allows the default value to be set from an
// environment variable named in the format of FLAG_NAME.
//
// This function panics if no default value is provided and the environment
// variable is invalid.
func Float64Var(p *float64, name string, value *float64, usage string) {
	switch s := os.Getenv(flagToEnv(name)); s {
	case "":
		if value != nil {
			flag.Float64Var(p, name, *value, flagUsage(name, usage))
			return
		}
	default:
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			flag.Float64Var(p, name, f, flagUsage(name, usage))
			return
		}
	}
	panic("invalid default float64 value for env " + flagToEnv(name))
}

// Just like flag.Int64Var, but allows the default value to be set from an
// environment variable named in the format of FLAG_NAME.
//
// This function panics if no default value is provided and the environment
// variable is invalid.
func Int64Var(p *int64, name string, value *int64, usage string) {
	switch s := os.Getenv(flagToEnv(name)); s {
	case "":
		if value != nil {
			flag.Int64Var(p, name, *value, flagUsage(name, usage))
			return
		}
	default:
		if i, err := strconv.ParseInt(s, 0, 64); err == nil {
			flag.Int64Var(p, name, i, flagUsage(name, usage))
			return
		}
	}
	panic("invalid default int64 value for env " + flagToEnv(name))
}

// Just like flag.IntVar, but allows the default value to be set from an
// environment variable named in the format of FLAG_NAME.
//
// This function panics if no default value is provided and the environment
// variable is invalid.
func IntVar(p *int, name string, value *int, usage string) {
	switch s := os.Getenv(flagToEnv(name)); s {
	case "":
		if value != nil {
			flag.IntVar(p, name, *value, flagUsage(name, usage))
			return
		}
	default:
		if i, err := strconv.ParseInt(s, 0, 0); err == nil {
			flag.IntVar(p, name, int(i), flagUsage(name, usage))
			return
		}
	}
	panic("invalid default int value for env " + flagToEnv(name))
}

// Just like flag.StringVar, but allows the default value to be set from an
// environment variable named in the format of FLAG_NAME.
//
// This function panics if no default value is provided and the environment
// variable is invalid.
func StringVar(p *string, name string, value *string, usage string) {
	switch s := os.Getenv(flagToEnv(name)); s {
	case "":
		if value != nil {
			flag.StringVar(p, name, *value, flagUsage(name, usage))
			return
		}
	default:
		flag.StringVar(p, name, s, flagUsage(name, usage))
		return
	}
	panic("invalid default string value for env " + flagToEnv(name))
}

// Just like flag.Uint64Var, but allows the default value to be set from an
// environment variable named in the format of FLAG_NAME.
//
// This function panics if no default value is provided and the environment
// variable is invalid.
func Uint64Var(p *uint64, name string, value *uint64, usage string) {
	switch s := os.Getenv(flagToEnv(name)); s {
	case "":
		if value != nil {
			flag.Uint64Var(p, name, *value, flagUsage(name, usage))
			return
		}
	default:
		if i, err := strconv.ParseUint(s, 0, 64); err == nil {
			flag.Uint64Var(p, name, i, flagUsage(name, usage))
			return
		}
	}
	panic("invalid default uint64 value for env " + flagToEnv(name))
}

// Just like flag.UintVar, but allows the default value to be set from an
// environment variable named in the format of FLAG_NAME.
//
// This function panics if no default value is provided and the environment
// variable is invalid.
func UintVar(p *uint, name string, value *uint, usage string) {
	switch s := os.Getenv(flagToEnv(name)); s {
	case "":
		if value != nil {
			flag.UintVar(p, name, *value, flagUsage(name, usage))
			return
		}
	default:
		if i, err := strconv.ParseUint(s, 0, 0); err == nil {
			flag.UintVar(p, name, uint(i), flagUsage(name, usage))
			return
		}
	}
	panic("invalid default uint value for env " + flagToEnv(name))
}
