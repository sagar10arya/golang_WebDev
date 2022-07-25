package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen(":8080", "tcp")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		io.WriteString(conn, "I see you connected")

		conn.Close()
	}
}
