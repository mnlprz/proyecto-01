package main

import (
	"github/mnlprz/go/proyecto-01/handlers"
	"log"
	"net/http"
)

func main() {

	handlers.SetHandlers()

	err := http.ListenAndServe(":5555", nil)
	if err != nil {
		log.Fatal(err)
	}

}
