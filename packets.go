package tflow

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"unsafe"

	"github.com/fatih/structs"
)

type Packet []interface{}

func FormatPacket(id int, payload interface{}) Packet {
	if structs.IsStruct(payload) {
		pl := structs.Values(payload)
		size := 3
		for i := 0; i < len(pl); i++ {
			size += int(unsafe.Sizeof(pl[i]))
		}
		buf := make(Packet, 2)
		buf[0] = (3 + size) << 3
		buf[1] = id
		return append(buf, pl...)
	}
	buf := make(Packet, 3)
	buf[0] = (3 + unsafe.Sizeof(payload)) << 3
	buf[1] = id
	buf[2] = payload
	return buf
}

func BuildPacket(p Packet) []byte {
	buf := new(bytes.Buffer)
	for i := 0; i < len(p); i++ {
		switch reflect.ValueOf(p[i]).Kind() {
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			b := new(bytes.Buffer)
			binary.Write(b, binary.LittleEndian, p[i])
			b.WriteTo(buf)
		case reflect.String:
			b := EncodeString(p[i].(string))
			buf.Write(b)
		}
	}
	return buf.Bytes()
}
