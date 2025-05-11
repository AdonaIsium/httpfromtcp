package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	msg, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("%v", err)
	}

	reader := io.Reader(msg)

	buf := make([]byte, 8)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			fmt.Printf("read: %s\n", buf[:n])
		}
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
	}

}
