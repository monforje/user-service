package main

import (
	"log"

	"github.com/monforje/user-service/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
