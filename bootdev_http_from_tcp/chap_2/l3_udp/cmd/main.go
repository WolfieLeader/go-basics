package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		defer f.Close()
		var sb strings.Builder
		buf := make([]byte, 8)
		for {
			n, err := f.Read(buf)
			if n > 0 {
				chunk := buf[:n]
				parts := bytes.Split(chunk, []byte{'\n'})

				for i := 0; i < len(parts)-1; i++ {
					sb.Write(parts[i])
					out <- sb.String()
					sb.Reset()
				}

				sb.Write(parts[len(parts)-1])
			}
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				log.Fatalf("Error reading file: %v", err)
			}
		}
		if sb.Len() > 0 {
			out <- sb.String()
		}
	}()
	return out
}

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("Error connecting to UDP: %v", err)
	}
	defer conn.Close()

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := r.ReadString('\n')
		if len(input) > 0 {
			_, writeErr := conn.Write([]byte(input))
			if writeErr != nil {
				log.Fatalf("Error writing to UDP: %v", writeErr)
			}
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("Error reading from stdin: %v", err)
		}
	}
}
