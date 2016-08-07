package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var addr = flag.String("addr", ":3030", "Address to bind this server to in the format server:port")

func main() {
	fmt.Println("Starting server on ", *addr)

	laddr, err := net.ResolveTCPAddr("tcp", *addr)

	if err != nil {
		log.Fatal("Failed to ResolveTCPAddr: ", err)
	}

	l, err := net.ListenTCP("tcp", laddr)

	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}

	defer l.Close()

	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// status, err := bufio.NewReader(conn).ReadString('\n')

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		if text != "::EOF::" {
			fmt.Println(text) // Println will add back the final '\n'
		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Fprintf(conn, "Message Received\n")

	conn.Close()
}
