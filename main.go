package main

import (
	"log"
	"os"

	"github.com/johnny-morrice/gosnake/snake"
)

func main() {
	app, err := snake.Setup()
	if err != nil {
		log.Printf("initialization failed: %s", err)
		os.Exit(1)
	}
	app.Run()
}
