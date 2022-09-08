package main

import (
	"context"
	"log"
)

func main() {
	server, err := buildWebServer(context.TODO())
	if err != nil {
		log.Panicln("failed to build")
	}
	server.Run(":8080")
}
