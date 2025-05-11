package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	resolver, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to resolve local addr: %v\n", err)
	}

	conn, err := net.DialUDP("udp", nil, resolver)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ListenUDP error: %v\n", err)
	}
	defer conn.Close()

	userInput := os.Stdin

	reader := bufio.NewReader(userInput)

	for {
		fmt.Printf(">")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("error reading user input: %v", err)
			continue
		}
		result, err := conn.Write([]byte(input))
		if err != nil {
			log.Printf("error writing to UDP: %v", err)
			continue
		}
		fmt.Printf("Sent %d bytes", result)
	}
}
