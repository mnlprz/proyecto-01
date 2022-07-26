package handlers

import (
	"github/mnlprz/go/proyecto-01/services"
	"log"
	"net/http"
)

func SetHandlers() {

	http.HandleFunc("/cargatabla", func(w http.ResponseWriter, req *http.Request) {
		err := services.CargaTabla()
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte("Tabla cargada exitosamente."))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := w.Write([]byte("HOLA"))
		if err != nil {
			log.Fatal(err)
		}
	})
}
