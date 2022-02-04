package envflag

import (
	"flag"
	"os"
	"strings"
	"time"
)

type EnvFlag struct {
	fs *flag.FlagSet
}

func New(name string, handler flag.ErrorHandling) *EnvFlag {
	return &EnvFlag{
		fs: flag.NewFlagSet(name, handler),
	}
}

func (ev *EnvFlag) Parse(args ...string) error {
	err := ev.fs.Parse(args)
	if err != nil {
		return err
	}
	ev.resetAsEnv()
	return nil
}

func (ev *EnvFlag) rebuildName(name string) string {
	return strings.ToUpper(strings.NewReplacer("-", "_", ".", "_", "@", "_", "#", "_").Replace(name))
}

func (ev *EnvFlag) resetAsEnv() {
	//build env value
	ev.fs.VisitAll(func(f *flag.Flag) {
		value, ok := os.LookupEnv(ev.rebuildName(f.Name))
		if !ok {
			return
		}
		f.Value.Set(value)
	})
}

func (ev *EnvFlag) String(name string, value string, usage string) *string {
	return ev.fs.String(name, value, usage)
}

func (ev *EnvFlag) Float64(name string, value float64, usage string) *float64 {
	return ev.fs.Float64(name, value, usage)
}

func (ev *EnvFlag) Duration(name string, value time.Duration, usage string) *time.Duration {
	return ev.fs.Duration(name, value, usage)
}

func (ev *EnvFlag) Uint64(name string, value uint64, usage string) *uint64 {
	return ev.fs.Uint64(name, value, usage)
}

func (ev *EnvFlag) Uint(name string, value uint, usage string) *uint {
	return ev.fs.Uint(name, value, usage)
}

func (ev *EnvFlag) Int64(name string, value int64, usage string) *int64 {
	return ev.fs.Int64(name, value, usage)
}

func (ev *EnvFlag) Int(name string, value int, usage string) *int {
	return ev.fs.Int(name, value, usage)
}

func (ev *EnvFlag) Bool(name string, value bool, usage string) *bool {
	return ev.fs.Bool(name, value, usage)
}

func (ev *EnvFlag) ProtoString(name string, value string, usage string) *string {
	out := new(string)
	ev.fs.Var(newProtoString(value, out), name, usage)
	return out
}

func (ev *EnvFlag) ProtoBytes(name string, value []byte, usage string) *[]byte {
	out := new([]byte)
	ev.fs.Var(newProtoBytes(value, out), name, usage)
	return out
}

func String(name string, value string, usage string) *string {
	return DefaultParser.String(name, value, usage)
}

func Float64(name string, value float64, usage string) *float64 {
	return DefaultParser.Float64(name, value, usage)
}

func Duration(name string, value time.Duration, usage string) *time.Duration {
	return DefaultParser.Duration(name, value, usage)
}

func Uint64(name string, value uint64, usage string) *uint64 {
	return DefaultParser.Uint64(name, value, usage)
}

func Uint(name string, value uint, usage string) *uint {
	return DefaultParser.Uint(name, value, usage)
}

func Int64(name string, value int64, usage string) *int64 {
	return DefaultParser.Int64(name, value, usage)
}

func Int(name string, value int, usage string) *int {
	return DefaultParser.Int(name, value, usage)
}

func Bool(name string, value bool, usage string) *bool {
	return DefaultParser.Bool(name, value, usage)
}

func ProtoString(name string, value string, usage string) *string {
	return DefaultParser.ProtoString(name, value, usage)
}

func Raw() *flag.FlagSet {
	return DefaultParser.Raw()
}

func (ev *EnvFlag) Raw() *flag.FlagSet {
	return ev.fs
}

func (ev *EnvFlag) Parsed() bool {
	return ev.fs.Parsed()
}

var DefaultParser = New("default_env_parser", flag.PanicOnError)

func Parsed() bool {
	return DefaultParser.Parsed()
}

func Parse() error {
	return DefaultParser.Parse(os.Args[1:]...)
}
