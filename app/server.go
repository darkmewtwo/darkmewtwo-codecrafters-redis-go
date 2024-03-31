package main

import (
	"fmt"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("+PONG\r\n"))
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()
	conn, err := l.Accept()
	for {
		// defer conn.Close()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		// fmt.Println(conn)
		buffer := make([]byte, 1024)
		buffN, _ := conn.Read(buffer)

		// fmt.Println(buffN)
		request := string(buffer[:buffN])
		fmt.Println(request)
		go handleConnection(conn)
		// fmt.Println(buffer)
		// cmd := strings.TrimSpace(request)
		// cmd_parts := strings.Split(cmd, " ")
		// fmt.Println(request, cmd, cmd_parts, (strings.TrimSpace(request)), len(request))
		// // fmt.Println(request, reflect.TypeOf(request), (strings.TrimSpace(request) == "PING"), len(strings.TrimSpace(request)), len(request))
		// if strings.ToLower(cmd_parts[2]) == "ping" {
		// 	conn.Write([]byte("+PONG\r\n"))
		// } else {
		// 	conn.Write([]byte("+Unrecognized cmd\r\n"))
		// }
	}
}
