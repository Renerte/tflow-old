package tflow

import (
	"bytes"
	"encoding/binary"
)

func PutVarint(val int, size uint) []byte {
	buf := make([]byte, 8)
	binary.PutVarint(buf, int64(val))
	if size > 0 {
		return buf[:size]
	}
	return bytes.Trim(buf, "\x00")
}

func EncodeString(text string) []byte {
	buf := PutVarint(len(text), 0)
	return append(buf, []byte(text)...)
}
