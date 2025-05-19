package main

import (
	"fmt"
	"log"
	"net"

	i "github.com/AdonaIsium/httpfromtcp/internal/request"
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
			req, err := i.RequestFromReader(c)
			if err != nil {
				fmt.Printf("received error: %v", err)
			}
			fmt.Printf("Request line:\n- Method: %s\n- Target: %s\n- Version: %s\nHeaders:\n", req.RequestLine.Method, req.RequestLine.RequestTarget, req.RequestLine.HttpVersion)
			for k, v := range req.Headers {
				fmt.Printf("- %s: %s\n", k, v)
			}
			fmt.Printf("Body:\n%s", string(req.Body))

		}(conn)

	}

}
