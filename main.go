package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	port = ":7070" // Gopher protocol typically uses port 70
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	_, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading from connection:", err)
		return
	}

	// Write the Gopher response
	response := "iWelcome to rcy's gopher site!  Buy more stuff!\t\terrormsg\r\n.\r\n"
	fmt.Fprint(conn, response)
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	log.Printf("Gopher server listening on port %s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
