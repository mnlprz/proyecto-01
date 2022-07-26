package main

import (
	"github/mnlprz/go/proyecto-01/handlers"
	"log"
	"net/http"
)

func main() {

	const PORT = ":5555"

	handlers.SetHandlers()

	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}
