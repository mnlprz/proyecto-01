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

	http.HandleFunc("/getentrada/{id}", func(w http.ResponseWriter, req *http.Request) {
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
		w.Write(entradaJson)
	})
}
