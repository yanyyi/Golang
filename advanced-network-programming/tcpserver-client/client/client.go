package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server at localhost on port 8080.
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Send a message to the server.
	conn.Write([]byte("Hello, server!"))
	buffer := make([]byte, 1024)
	// Read the response from the server.
	conn.Read(buffer)
	fmt.Println(string(buffer))
}
