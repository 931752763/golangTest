package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	b := isPortListening("tcp", "localhost", 8091)
	fmt.Println(b)
}

//
func isPortListening(protocol string, hostname string, port int) bool {
	fmt.Printf("scanning port %d \n", port)
	p := strconv.Itoa(port)
	addr := net.JoinHostPort(hostname, p)
	conn, err := net.DialTimeout(protocol, addr, 3*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
