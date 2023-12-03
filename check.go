package main

import (
	"fmt"
	"net"
	"time"
)

// Check checks the reachability of a destination address and port.
// It takes the destination address and port as input parameters.
// It returns a string indicating the status of the reachability.
// If the destination is unreachable, it returns a string in the format "[DOWN] <destination> is unreachable, Error: <error>".
// If the destination is reachable, it returns a string in the format "[UP] <destination> is reachable, From: <local address>, To: <remote address>".
func Check(destination string, port string) string {
	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, \n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable, \n From: %v\n To: %v", destination,
			conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}
