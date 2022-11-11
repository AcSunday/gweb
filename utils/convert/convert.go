package convert

import (
	"bytes"
	"strings"
)

func MustBytesToString(bs []byte) string {
	return bytes.NewBuffer(bs).String()
}

func MustStringToBytes(str string) (ret []byte) {
	_, err := strings.NewReader(str).Read(ret)
	if err != nil {
		panic(any(err))
	}
	return
}
