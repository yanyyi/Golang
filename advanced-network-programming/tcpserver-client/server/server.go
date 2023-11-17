package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen on TCP port 8080 on all available unicast and
	// any unicast IP addresses.
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	// Infinite loop to handle incoming connections
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Launch a new goroutine to handle the connection
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Send a response back to the client.
	conn.Write([]byte("Received: " + string(buffer)))
}
