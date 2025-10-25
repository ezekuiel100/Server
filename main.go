package main

import (
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", ":8000")
	defer ln.Close()

	fmt.Println("Servidor escutando em :8080...")

	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	fmt.Println("Nova conexão de:", conn)

}
