package envflag

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	spliter = "://"
)

type protoString string
type protoBytes []byte

func newProtoString(value string, p *string) *protoString {
	*p = value
	return (*protoString)(p)
}

func newProtoBytes(value []byte, p *[]byte) *protoBytes {
	*p = value
	return (*protoBytes)(p)
}

func (s *protoString) Set(str string) error {
	value, err := defaultProtoDec.DecodeString(str)
	if err != nil {
		return err
	}
	*s = protoString(value)
	return nil
}

func (s *protoString) String() string {
	return string(*s)
}

func (b *protoBytes) Set(s string) error {
	value, err := defaultProtoDec.DecodeString(s)
	if err != nil {
		return err
	}
	*b = value
	return nil
}

func (s *protoBytes) String() string {
	return hex.EncodeToString(*s)
}

var defaultProtoDec = &protoDec{}

type protoDec struct {
}

func (d *protoDec) DecodeString(str string) ([]byte, error) {
	var idx int
	if idx = strings.Index(str, spliter); idx < 0 {
		return []byte(str), nil
	}
	proto := str[:idx]
	value := str[idx+len(spliter):]
	return d.DecodeProtoString(proto, value)
}

func (d *protoDec) DecodeProtoString(proto string, str string) ([]byte, error) {
	//也没几种协议, 直接switch case吧
	switch strings.ToLower(proto) {
	case "direct":
		return d.decodeProtoDirect(str)
	case "base64":
		return d.decodeProtoBase64(str)
	case "hex":
		return d.decodeProtoHex(str)
	default:
		return d.decodeProtoUnknow(str)
	}
}

func (d *protoDec) decodeProtoBase64(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func (d *protoDec) decodeProtoHex(str string) ([]byte, error) {
	return hex.DecodeString(str)
}

func (d *protoDec) decodeProtoDirect(str string) ([]byte, error) {
	return []byte(str), nil
}

func (d *protoDec) decodeProtoUnknow(str string) ([]byte, error) {
	return nil, fmt.Errorf("unknow proto")
}
