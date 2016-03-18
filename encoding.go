package tflow

import (
	"bytes"
	"encoding/binary"
)

func PutUvarint(val int, size uint) []byte {
	buf := make([]byte, 8)
	binary.PutUvarint(buf, uint64(val))
	if size > 0 {
		return buf[:size]
	}
	return bytes.Trim(buf, "\x00")
}

func EncodeString(text string) []byte {
	buf := PutUvarint(len(text), 0)
	return append(buf, []byte(text)...)
}

func DecodeString(raw []byte) string {
	r, _ := binary.Uvarint(raw)
	return string(raw[len(raw)-int(r)-1 : len(raw)-(len(raw)-int(r))+1])
}
