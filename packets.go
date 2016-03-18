package tflow

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"

	"github.com/fatih/structs"
)

type Packet []interface{}

func BuildPacket(id byte, payload interface{}) Packet {
	if structs.IsStruct(payload) {
		pl := structs.Values(payload)
		buf := make(Packet, 1)
		buf[0] = id
		return append(buf, pl...)
	}
	buf := make(Packet, 2)
	buf[0] = id
	buf[1] = payload
	return buf
}

func FormatPacket(p Packet) []byte {
	buf := new(bytes.Buffer)
	for i := 0; i < len(p); i++ {
		switch reflect.ValueOf(p[i]).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			binary.Write(buf, binary.LittleEndian, p[i])
		case reflect.String:
			buf.Write(EncodeString(p[i].(string)))
		default:
			fmt.Printf("Did not recognise %v (its type is %v)", p[i], reflect.TypeOf(p[i]).Name())
		}
	}
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, int16(len(buf.Bytes()))+2)
	buf.WriteTo(b)
	fmt.Println(b.Bytes())
	return b.Bytes()
}

func ParsePacket(r io.Reader) {
	lb := make([]byte, 2)
	r.Read(lb)
	lbr := bytes.NewReader(lb)
	var l int16
	binary.Read(lbr, binary.LittleEndian, &l)
	fmt.Printf("Length of packet: %v\n", l)
	id := make([]byte, 1)
	r.Read(id)
	fmt.Printf("Packet id: %v\n", id)
	raw := make([]byte, l-2)
	r.Read(raw)
	PacketHandlers[id[0]](raw)
}
