package main

import (
	"log"
	"net/http"

	"github.com/mnlprz/go/proyecto-01/handlers"
)

func main() {

	const PORT = ":5555"

	handlers.SetHandlers()

	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}
