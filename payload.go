package tflow

import (
	"errors"
	"fmt"
)

type PacketHandler func([]byte) error

var PacketHandlers = map[byte]PacketHandler{1: ConnectHandler, 2: DisconnectHandler, 3: ContinueConnectingHandler}

func ConnectHandler(raw []byte) error {
	if DecodeString(raw) != "Terraria156" {
		return errors.New("Incorrect Terraria version!")
	}
	return nil
}

func DisconnectHandler(raw []byte) error {
	fmt.Println(DecodeString(raw))
	return nil
}

func ContinueConnectingHandler(raw []byte) error {
	fmt.Printf("User id: %v", raw[0])
	return nil
}
