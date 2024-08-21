package main

// interfaces: Lists all network interfaces on the server.
// ip: Returns the server's IP address.
// time: Returns the server's local time.
// rtt: Measures round-trip time for a 1024-byte message in ms.
// speed: Measures transmission speed for an 8192-byte message in kB/s.

import (
	cmd "github.com/arkantrust/wadjet/cli/commands"
)

func main() {
	cmd.Execute()
}
