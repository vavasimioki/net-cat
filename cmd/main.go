package main

import (
	"log"
	"netcat/server"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		s := server.NewServer(":8989")
		log.Fatal(s.Start())
	} else if len(os.Args) == 2 {
		port := os.Args[1]
		s := server.NewServer(":" + port)

		log.Fatal(s.Start())
	} else {
		log.Fatal("[USAGE]: ./TCPChat $port")
	}
}
