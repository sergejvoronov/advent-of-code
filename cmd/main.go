package main

import (
	"log"

	"github.com/sergejvoronov/advent-of-code/internal/app"
)

func main() {
	if err := app.Exec(); err != nil {
		log.Fatal(err)
	}
}
