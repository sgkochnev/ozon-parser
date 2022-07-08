package main

import (
	"log"
	"os"
	"ozon-parser/internal/app"
)

func main() {
	// cfg := config.Get()

	if err := app.Run(); err != nil {
		log.Println("error")
		os.Exit(1)
	}

}
