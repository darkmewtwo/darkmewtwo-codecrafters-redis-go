package main

import (
	"fmt"
	"strings"

	// Uncomment this block to pass the first stage

	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	// defer conn.Close()
	// fmt.Println(conn)
	for {

		buffer := make([]byte, 1024)
		buffN, _ := conn.Read(buffer)

		// fmt.Println(buffN)
		request := string(buffer[:buffN])
		fmt.Println(request)
		cmd := strings.TrimSpace(request)
		cmd_parts := strings.Split(cmd, "\\r\\n")
		// fmt.Println(cmd, cmd_parts)
		// log.Println(cmd, cmd_parts)
		keyword := ""
		if len(cmd_parts) > 1 {
			keyword = strings.ToLower(cmd_parts[2])
		}
		fmt.Println(keyword)
		switch keyword {

		case "ping":
			conn.Write([]byte("+PONG\r\n"))
		case "redis-cli ping":
			conn.Write([]byte("+PONG\r\n"))
		case "echo":
			message := "+" + cmd_parts[4]
			conn.Write([]byte(message))
		default:
			conn.Write([]byte("+PONG\r\n"))
		}

	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
	// Uncomment this block to pass the first stage

	//

	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		// defer conn.Close()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}
