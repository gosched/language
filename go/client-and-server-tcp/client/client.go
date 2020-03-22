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
	addr = flag.String("addr", "127.0.0.1:8972", "server address")
)

// go run client.go -addr=127.0.0.1:8972
// write
// write
// close_conn
// exit

func main() {
	flag.Parse()

	conn, err := net.Dial("tcp", *addr)
	panicOnErr(err)

	go func() {
		var buf = make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				log.Printf("read err from server %s: %v", *addr, err)
				return
			}
			log.Printf("read %d bytes from server %s", n, *addr)
		}
	}()

	id := 0
	write := func() {
		msg := fmt.Sprintf("sent clientid: %d from client", id)
		id++
		n, err := conn.Write([]byte(msg))
		if err != nil {
			log.Printf("write err to server %s: %v", *addr, err)
			return
		}
		log.Printf("write %d bytes to server %s", n, *addr)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := scanner.Text()
		switch cmd {
		case "close_conn":
			conn.Close()
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
