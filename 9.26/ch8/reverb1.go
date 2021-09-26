package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintf(c, "\n", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintf(c, "\n", shout)
	time.Sleep(delay)
	fmt.Fprintf(c, "\n", strings.ToLower(shout))

}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}
