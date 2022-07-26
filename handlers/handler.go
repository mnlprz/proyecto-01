package handlers

import (
	"github/mnlprz/go/proyecto-01/utils"
	"log"
	"net/http"
)

func SetHandlers() {

	http.HandleFunc("/cargatabla", func(w http.ResponseWriter, req *http.Request) {
		err := utils.CargaTable()
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := w.Write([]byte("HOLA"))
		if err != nil {
			log.Fatal(err)
		}
	})
}
