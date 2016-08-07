package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var serverAddr = flag.String("serverAddr", ":3030", "Server address to connect to in the format server:port")

func main() {
	for {
		fmt.Printf("Connecting to server %v\n", *serverAddr)

		conn, err := net.Dial("tcp", *serverAddr)
		if err != nil {
			fmt.Println("dial error:", err)
			return
		}
		defer func() {
			if err = conn.Close(); err != nil {
				fmt.Printf("Failed to close connection: %v", err)
				return
			}

			fmt.Println("Connection Closed")
		}()

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Send a Message: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading user input: ", err)
			break
		}

		// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
		fmt.Fprintf(conn, text+"\n")
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("read error: ", err)
			break
		}

		fmt.Printf("Response received: %v\n", response)

	}

}
