package handlers

import (
	"encoding/json"
	"github/mnlprz/go/proyecto-01/models"
	"github/mnlprz/go/proyecto-01/services"
	"log"
	"net/http"
	"strconv"
)

func SetHandlers() {

	http.HandleFunc("/cargatablaofertas", func(w http.ResponseWriter, req *http.Request) {
		err := services.CargaTablaOfertas()
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte("Tabla OFERTAS cargada exitosamente."))
	})

	http.HandleFunc("/cargatablapersonas", func(w http.ResponseWriter, req *http.Request) {
		err := services.CargaTablaPersonas()
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte("Tabla PERSONAS cargada exitosamente."))
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
			var entrada models.Entrada
			dec := json.NewDecoder(req.Body)
			err := dec.Decode(&entrada)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(entrada)
			err = services.InsertEntrada(entrada)
			if err != nil {
				log.Fatal(err)
			}
			w.Write([]byte("Insertado correctamente."))

		case "PUT":
			var entrada models.Entrada
			dec := json.NewDecoder(req.Body)
			err := dec.Decode(&entrada)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(entrada)
			err = services.UpdateEntrada(entrada)
			if err != nil {
				log.Fatal(err)
			}
			w.Write([]byte("Actualizado correctamente."))

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
