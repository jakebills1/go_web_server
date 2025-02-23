package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	logger = log.New(os.Stdout, "logger: ", log.Lshortfile)
)

func handleInterrupt(c chan os.Signal) {
	<-c
	logger.Println("exiting due to sigint")
	os.Exit(1)
}
func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go handleInterrupt(c)
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
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		logger.Println("got message from conn:", scanner.Text())
		bytesWritten, err := conn.Write(append(scanner.Bytes(), '\n'))
		if err != nil {
			logger.Fatal("Write():", err)
		}
		logger.Println("wrote", bytesWritten, "bytes to conn")
	}
	logger.Println("closing connection")
}
