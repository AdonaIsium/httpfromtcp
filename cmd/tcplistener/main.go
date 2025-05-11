package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	reader := io.Reader(f)
	buf := make([]byte, 8)
	currentLine := ""
	c := make(chan string)

	go func() {
		defer close(c)
		defer f.Close()

		for {
			n, err := reader.Read(buf)
			if n > 0 {
				chunk := string(buf[:n])
				lines := strings.Split(chunk, "\n")
				for i, line := range lines {
					if i == 0 {
						currentLine += line
					} else {
						c <- currentLine
						currentLine = line
					}
				}
			}
			if err == io.EOF {
				if currentLine != "" {
					c <- currentLine
				}
				break
			}

			if err != nil {
				log.Fatal(err)
			}
		}

	}()
	return c
}

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
			msgChan := getLinesChannel(c)
			for msg := range msgChan {
				fmt.Printf("%s\n", msg)
			}
			c.Close()

		}(conn)

	}

}
