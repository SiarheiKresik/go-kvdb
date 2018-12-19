package main

import (
	"flag"
	"log"
	"net"
	"strconv"
)

const protocol = "tcp"

type command struct {
	fields []string
	result chan string
}

func main() {

	flag.Parse()

	addr := net.JoinHostPort("", strconv.Itoa(config.port))
	li, err := net.Listen(protocol, addr)
	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()

	log.Printf("Server is listening %s\n", addr)
	log.Println("Ready to accept connections")

	commands := make(chan command)

	go storage(commands, config.mode)

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(commands, conn)
	}
}
