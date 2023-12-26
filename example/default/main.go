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
	envflag.BoolVar(&f.myBool, "my-bool", envflag.Ptr(true), "a bool value")
	envflag.DurationVar(&f.myDuration, "my-duration", envflag.Ptr(time.Minute), "a duration value")
	envflag.Float64Var(&f.myFloat64, "my-float64", envflag.Ptr(1.1), "a float64 value")
	envflag.Int64Var(&f.myInt64, "my-int64", envflag.Ptr(int64(1)), "a int64 value")
	envflag.IntVar(&f.myInt, "my-int", envflag.Ptr(2), "a int value")
	envflag.StringVar(&f.myString, "my-string", envflag.Ptr("str"), "a string value")
	envflag.TextVar(&f.myText, "my-text", LogLevel("debug"), "a text value of either 'debug', 'info', or 'error'")
	envflag.Uint64Var(&f.myUint64, "my-uint64", envflag.Ptr(uint64(3)), "a uint64 value")
	envflag.UintVar(&f.myUint, "my-uint", envflag.Ptr(uint(4)), "a uint value")
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
