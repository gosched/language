package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	addr = flag.String("addr", ":8972", "listened address")
)

// go run server.go
// write
// write
// close_conn
// close_ln
// exit

func main() {
	flag.Parse()
	ln, err := net.Listen("tcp", *addr)
	panicOnErr(err)

	conn, err := ln.Accept()
	panicOnErr(err)
	clientAddr := conn.RemoteAddr().String()

	go func() {
		var buf = make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				log.Printf("read err from client %s: %v", clientAddr, err)
				return
			}
			log.Printf("read %d bytes from client %s", n, clientAddr)
		}
	}()

	id := 0
	write := func() {
		msg := fmt.Sprintf("sent id: %d from server", id)
		id++
		n, err := conn.Write([]byte(msg))
		if err != nil {
			log.Printf("write err to client %s: %v", clientAddr, err)
			return
		}
		log.Printf("write %d bytes to client %s", n, clientAddr)
	}

	go func() {
		for {
			_, err := ln.Accept()
			if err != nil {
				log.Printf("accept err : %v", err)
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := scanner.Text()
		switch cmd {
		case "close_conn":
			conn.Close()
		case "close_ln":
			ln.Close()
			// return
		case "write":
			write()
		case "exit", "quit":
			return
		}
	}
}
func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
