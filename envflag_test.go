package envflag

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	ev := New("hello", flag.PanicOnError)
	isOk := ev.Bool("is_ok", false, "")
	num66 := ev.Int64("num66", 0, "")
	float55 := ev.Float64("float55", 11.55, "")
	float77 := ev.Float64("float77", 11, "")
	float88 := ev.Float64("float88", 88, "")
	float99 := ev.Float64("float99", 0, "")

	os.Setenv("IS_OK", "true")
	os.Setenv("NUM66", "66")
	os.Setenv("FLOAT55", "55")
	if err := ev.Parse("--float77=77", "--float99=99"); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, *isOk)
	assert.Equal(t, int64(66), *num66)
	assert.Equal(t, float64(55), *float55)
	assert.Equal(t, float64(77), *float77)
	assert.Equal(t, float64(88), *float88)
	assert.Equal(t, float64(99), *float99)
}

func TestSpecName(t *testing.T) {
	ev := New("haha", flag.PanicOnError)
	val := ev.Int64("a-b", 123, "aaaa")
	os.Setenv("A_B", "6")
	ev.Parse("--a-b=5")
	assert.Equal(t, int64(6), *val)
}

func TestSpecName1(t *testing.T) {
	{
		ev := New("haha", flag.PanicOnError)
		val := ev.Int64("aa.bb", 123, "aaaa")
		ev.Parse("--aa.bb=5")
		assert.Equal(t, int64(5), *val)
	}
	{
		ev := New("haha", flag.PanicOnError)
		val := ev.Int64("c.d", 123, "aaaa")
		os.Setenv("C_D", "6")
		ev.Parse("--c.d=5")
		assert.Equal(t, int64(6), *val)
	}
}

func TestProtoBytes(t *testing.T) {
	{
		ev := New("haha", flag.PanicOnError)
		val := ev.ProtoBytes("byte", []byte("hello world"), "aaa")
		ev.Parse("--byte=base64://" + base64.StdEncoding.EncodeToString([]byte("this is a test")))
		assert.Equal(t, []byte("this is a test"), *val)
	}
	{
		ev := New("haha", flag.PanicOnError)
		val := ev.ProtoBytes("byte", []byte("hello world"), "aaa")
		assert.Equal(t, []byte("hello world"), *val)
	}
	{
		ev := New("haha", flag.PanicOnError)
		val := ev.ProtoBytes("byte", []byte("hello world"), "aaa")
		ev.Parse("--byte=hex://" + hex.EncodeToString([]byte("this is a test")))
		assert.Equal(t, []byte("this is a test"), *val)
	}
}

func TestProtoString(t *testing.T) {
	{
		ev := New("haha", flag.PanicOnError)
		val := ev.ProtoString("str", "hello world", "aaa")
		ev.Parse("--str=base64://" + base64.StdEncoding.EncodeToString([]byte("this is a test")))
		assert.Equal(t, "this is a test", *val)
	}
	{
		ev := New("haha", flag.PanicOnError)
		val := ev.ProtoString("str", "hello world", "aaa")
		assert.Equal(t, "hello world", *val)
	}
	{
		ev := New("haha", flag.PanicOnError)
		val := ev.ProtoString("str", "hello world", "aaa")
		ev.Parse("--str=hex://" + hex.EncodeToString([]byte("this is a test")))
		assert.Equal(t, "this is a test", *val)
	}
	{
		ev := New("haha", flag.PanicOnError)
		val := ev.ProtoString("str", "hello world", "aaa")
		ev.Parse("--str=direct://" + "this is a test")
		assert.Equal(t, "this is a test", *val)
	}
	{
		ev := New("haha", flag.PanicOnError)
		val := ev.ProtoString("str", "hello world", "aaa")
		os.Setenv("STR", "hex://"+hex.EncodeToString([]byte("this is a test2")))
		ev.Parse("--str=hex://" + hex.EncodeToString([]byte("this is a test")))
		assert.Equal(t, "this is a test2", *val)
	}
}
