package tflow

import (
	"flag"
	"log"
	"net"
)

type Server struct {
	MaxPlayers byte
}

func handleConnection(conn net.Conn) {
	log.Printf("Client %v", conn.LocalAddr())
	if err := ParsePacket(conn); err != nil {
		conn.Write(FormatPacket(BuildPacket(2, err)))
		return
	}
	conn.Write(FormatPacket(BuildPacket(2, "Server is not ready. This is a placeholder.")))
	conn.Close()
}

func StartServer(max byte) Server {
	flag.Parse()
	log.Printf("Test Server by Renerte")
	ln, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalf("Could not start server!: %v", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Could not start connection: %v", err)
		}
		go handleConnection(conn)
	}
}
