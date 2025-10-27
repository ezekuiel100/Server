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
	defer conn.Close()
	fmt.Println("ip de quem conectou:", conn.RemoteAddr())

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Erro de leitura:", err)
		return
	}

	fmt.Println("Mensagem:", string(buf[:n]))
	conn.Write([]byte("Recebido!\n"))

}

//
