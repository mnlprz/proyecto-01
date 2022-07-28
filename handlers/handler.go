package handlers

import (
	"encoding/json"
	"github/mnlprz/go/proyecto-01/services"
	"log"
	"net/http"
	"strconv"
)

func SetHandlers() {

	http.HandleFunc("/cargatabla", func(w http.ResponseWriter, req *http.Request) {
		err := services.CargaTabla()
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte("Tabla cargada exitosamente."))
	})

	http.HandleFunc("/entrada/{id}", func(w http.ResponseWriter, req *http.Request) {

		switch req.Method {

		case "GET":
			idParam := req.URL.Query().Get("id")
			id, err := strconv.ParseInt(idParam, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			entrada, err := services.GetEntrada(id)
			if err != nil {
				log.Fatal(err)
			}
			entradaJson, err := json.Marshal(entrada)
			if err != nil {
				log.Fatal(err)
			}
			w.Write(entradaJson)

		case "POST":
			log.Println("METODO POST")

		case "PUT":
			log.Println("METODO PUT")

		case "DELETE":
			idParam := req.URL.Query().Get("id")
			id, err := strconv.ParseInt(idParam, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			err = services.DeleteEntrada(id)
			if err != nil {
				log.Fatal(err)
			}
			w.Write([]byte("Eliminado correctamente."))
		}
	})

	http.HandleFunc("/getentradas/{campo1}", func(w http.ResponseWriter, req *http.Request) {
		campo1 := req.URL.Query().Get("campo1")
		entradas, err := services.GetEntradas(campo1)
		if err != nil {
			log.Fatal(err)
		}
		entradaJson, err := json.Marshal(entradas)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(entradaJson)
	})
}
