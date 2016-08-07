package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

var addr = flag.String("addr", ":3030", "Address to bind this server to in the format server:port")

func main() {
	fmt.Println("Starting server on ", *addr)

	l, err := net.Listen("tcp", *addr)

	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

	defer l.Close()

	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Connection error: ", err)
			break
		}

		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	// Shut down the connection.
	defer func() {
		if err := c.Close(); err != nil {
			fmt.Printf("Failed to close connection: %v", err)
			return
		}

		fmt.Println("Connection Closed")
	}()

	fmt.Println("Handling connection...")
	if message, err := bufio.NewReader(c).ReadString('\n'); err != nil && err != io.EOF {
		fmt.Println("read error:", err)
	} else {
		fmt.Printf("Message received: %s", message)
		c.Write([]byte(message + "\n"))
	}

	// Echo all incoming data.
	// io.Copy(c, c)

	// c.Write([]byte("thanks!\n"))

}
