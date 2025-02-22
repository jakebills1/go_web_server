package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

var (
	logger = log.New(os.Stdout, "logger: ", log.Lshortfile)
)

func main() {
	ln, err := net.Listen("tcp", ":2000")

	if err != nil {
		log.Fatalf("Listen(): %v", err)
	}

	logger.Println("waiting for new connections")
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Fatalf("Accept(): %v", err)
		}
		logger.Println("accepted new connection")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		logger.Println("got message from conn:", scanner.Text())
		bytesWritten, err := conn.Write(append(scanner.Bytes(), '\n'))
		if err != nil {
			logger.Fatal("Write():", err)
		}
		logger.Println("wrote", bytesWritten, "bytes to conn")
	}
	err := conn.Close()
	if err != nil {
		logger.Fatal("Close():", err)
	}
}
