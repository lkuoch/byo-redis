package main

import (
	"fmt"
	"net"
	"os"
)

const (
	listenAddr = "0.0.0.0:6379"
)

func writeResponse(c net.Conn, msg string) {
	_, err := c.Write([]byte(msg))

	if err != nil {
		fmt.Println("Error writing response: ", err.Error())
	}
}

func parseConnections(c net.Conn) {
	defer c.Close()

	buffer := make([]byte, 1024)
	n, err := c.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		return
	}

	request := string(buffer[:n])
	fmt.Println("Recieved request: ", request)

	// If request is `PING` respond with PONG
	if request == "PING\r\n" {
		writeResponse(c, "+PONG\r\n")
	} else {
		writeResponse(c, "???\r\n")
	}
}

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go parseConnections(c)
	}
}
