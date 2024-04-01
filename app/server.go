package main

import (
	"fmt"
	"strings"

	// Uncomment this block to pass the first stage

	"net"
	"os"
)

const (
	CRLF             = "\r\n"
	SIMPLE_STRING    = "+"
	SIMPLE_ERROR     = "-"
	INTEGER          = ":"
	BULK_STRINGS     = "$"
	ARRAY            = "*"
	BULK_NULL_STRING = "-1"
)

var dataStore map[string]string

func init() {
	dataStore = make(map[string]string)
}

func constructResponseMessage(resp_data_type string, data string) string {

	message := resp_data_type + data + CRLF
	return message
}

func setData(key string, value string) {
	dataStore[key] = value
	// return true
}

func getData(key string) (string, string) {
	value, exists := dataStore[key]
	if exists {
		return SIMPLE_STRING, value
	} else {
		return BULK_STRINGS, BULK_NULL_STRING
	}
}

func handleConnection(conn net.Conn) {
	// defer conn.Close()
	// fmt.Println(conn)
	for {

		buffer := make([]byte, 1024)
		buffN, _ := conn.Read(buffer)

		// fmt.Println(buffer)
		request := string(buffer[:buffN])
		fmt.Println(strings.Fields(request), request, "M")
		// cmd := strings.TrimSpace(request)
		// cmd_parts := strings.Split(cmd, "\\r\\n")
		cmd_parts := strings.Fields(request)
		// fmt.Println(cmd, cmd_parts)
		// log.Println(cmd, cmd_parts)
		keyword := ""
		if len(cmd_parts) > 1 {
			keyword = strings.ToLower(cmd_parts[2])
		}
		// fmt.Println(keyword)
		switch keyword {

		case "ping":
			conn.Write([]byte(constructResponseMessage(SIMPLE_STRING, "PONG")))
		case "echo":
			conn.Write([]byte(constructResponseMessage(SIMPLE_STRING, cmd_parts[4])))
		case "set":
			setData(cmd_parts[4], cmd_parts[6])
			conn.Write([]byte(constructResponseMessage(SIMPLE_STRING, "OK")))
		case "get":
			dataType, value := getData(cmd_parts[4])
			conn.Write([]byte(constructResponseMessage(dataType, value)))
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
