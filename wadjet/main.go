package main

// start a tcp server that listens on port 8080 and depending on the message received, it will respond with the output of a function
// that corresponds to the message.

import (
	"fmt"
	"net"
	"github.com/arkantrust/wadjet/cmd"
)

// handleConnection handles the connection by reading the message, calling the corresponding function and sending the response.
func handleConnection(conn net.Conn) {
	// read the request
	buf := make([]byte, 1024)
	n, nerr := conn.Read(buf)
	if nerr != nil {
		fmt.Println("Error reading:", nerr)
		return
	}
	req := string(buf[:n])

	// call the corresponding function
	var res string
	var err error = nil

	switch req {
	case "interfaces":
		res, err = cmd.Interfaces()
	case "ip":
		res, err = cmd.IP()
	case "time":
		res, err = cmd.Time()
	case "rtt":
		res, err = cmd.RTT()
	case "speed":
		res, err = cmd.Speed()
	default:
		res, err = "Invalid command.", nil
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// send the response
	_, err = conn.Write([]byte(res))
	if err != nil {
		fmt.Println("Error writing:", err)
	}
}

const PORT = 8080

func main() {
	// listen on port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	fmt.Printf("Listening on port %d...\n", PORT)
	defer ln.Close()

	// accept connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			return
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())
		defer conn.Close()

		// handle the connection
		go handleConnection(conn)
	}
}