# envflag

Helper functions wrapping the `*Var` functions of "flag" package, but default values are optional

```diff
- func DurationVar(p *time.Duration, name string, value time.Duration, usage string)
+ func DurationVar(p *time.Duration, name string, value *time.Duration, usage string)
```

and most importantly, default values can be provided by env. e.g. a flag named `session-duration` could have its default value set by env var `SESSION_DURATION`

# Panic

Command-line flags are things that we want to parse upfront and abort if the values are invalid.
So, though we don't like Go functions to `panic`, it's fine in this case.

- if `SESSION_DURATION=42m` is set, yes [42 * time.Minute](https://pkg.go.dev/time#ParseDuration) is the default value
- But if `SESSION_DURATION=42` is set by accident, the program should abort instead of defaulting to the hardcoded default of `5 * time.Hour`

# Usage

```go
type Flag struct {
	myBool     bool
	myDuration time.Duration
	myFloat64  float64
	myInt64    int64
	myInt      int
	myString   string
	myText     LogLevel // 'debug', 'info', or 'error'
	myUint64   uint64
	myUint     uint
}

func main() {
	f := Flag{}
	envflag.BoolVar(&f.myBool, "my-bool", envflag.Ptr(true), "a bool value")
	envflag.DurationVar(&f.myDuration, "my-duration", envflag.Ptr(time.Minute), "a duration value")
	envflag.Float64Var(&f.myFloat64, "my-float64", envflag.Ptr(1.1), "a float64 value")
	envflag.Int64Var(&f.myInt64, "my-int64", envflag.Ptr(int64(1)), "a int64 value")
	envflag.IntVar(&f.myInt, "my-int", nil, "a int value")
	envflag.StringVar(&f.myString, "my-string", nil, "a string value")
	envflag.TextVar(&f.myText, "my-text", LogLevel("debug"), "a text value of either 'debug', 'info', or 'error'")
	envflag.Uint64Var(&f.myUint64, "my-uint64", nil, "a uint64 value")
	envflag.UintVar(&f.myUint, "my-uint", nil, "a uint value")
	envflag.Parse() //  or the standard `flag.Parse()`

	fmt.Printf("%#v\n", f)
}
```