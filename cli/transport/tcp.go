package transport

import (
	"net"
)

// send sends a message to the server by specifying the server's address, port and the message to send.
// it creates the connection, sends the message, returns the response and closes the connection gracefully.
func Send(address string, port string, message string) (string, error) {
	// create the connection
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// send the message
	_, err = conn.Write([]byte(message))
	if err != nil {
		return "", err
	}

	// read the response
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}
