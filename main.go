package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Servidor escutando em :8000...")

	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println(conn.RemoteAddr())

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Erro de leitura:", err)
		return
	}

	fmt.Println(string(buf[:n]))

	response := "HTTP/1.1 200 OK\r\n\r\nHello, world!"
	conn.Write([]byte(response))
}
