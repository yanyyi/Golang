package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen on TCP port 8080.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		// Accept a connection.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Handle the connection in a new goroutine.
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	// Read the incoming connection.
	conn.Read(buffer)
	fmt.Println("Received:", string(buffer))
	// Respond to the client.
	conn.Write([]byte("Message received!"))
}
