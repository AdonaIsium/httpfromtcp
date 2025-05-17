package main

import (
	"fmt"
	"log"
	"net"

	i "github.com/AdonaIsium/httpfromtcp/internal"
)

func main() {
	l, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Printf("Error listening, too hard: %v", err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			i.RequestFromReader(c)
			msgChan := getLinesChannel(c)
			for msg := range msgChan {
				fmt.Printf("%s\n", msg)
			}
			c.Close()

		}(conn)

	}

}
