package main

import (
	"fmt"
	"time"

	"github.com/choonkeat/envflag-go"
)

type Flag struct {
	myBool     bool
	myDuration time.Duration
	myFloat64  float64
	myInt64    int64
	myInt      int
	myString   string
	myText     LogLevel
	myUint64   uint64
	myUint     uint
}

func main() {
	f := Flag{}
	envflag.BoolVar(&f.myBool, "my-bool", nil, "a bool value")
	envflag.DurationVar(&f.myDuration, "my-duration", nil, "a duration value")
	envflag.Float64Var(&f.myFloat64, "my-float64", nil, "a float64 value")
	envflag.Int64Var(&f.myInt64, "my-int64", nil, "a int64 value")
	envflag.IntVar(&f.myInt, "my-int", nil, "a int value")
	envflag.StringVar(&f.myString, "my-string", nil, "a string value")
	envflag.TextVar(&f.myText, "my-text", nil, "a text value of either 'alpha', 'beta', or 'charlie'")
	envflag.Uint64Var(&f.myUint64, "my-uint64", nil, "a uint64 value")
	envflag.UintVar(&f.myUint, "my-uint", nil, "a uint value")
	envflag.Parse() //  or the standard `flag.Parse()`

	fmt.Printf("%#v\n", f)
}

type LogLevel string

func (rs LogLevel) MarshalText() ([]byte, error) {
	switch string(rs) {
	case "debug", "info", "error":
		return []byte(rs), nil
	default:
		return nil, fmt.Errorf("value '%s' is not allowed", rs)
	}
}

func (rs *LogLevel) UnmarshalText(text []byte) error {
	switch string(string(text)) {
	case "debug", "info", "error":
		*rs = LogLevel(text)
		return nil
	default:
		return fmt.Errorf("value '%s' is not allowed", text)
	}
}
