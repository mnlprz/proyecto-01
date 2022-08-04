package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/mnlprz/go/proyecto-01/models"
	"github.com/mnlprz/go/proyecto-01/services"
)

type Contrato struct {
	Nup         string
	CodEntidad  string
	CodCentro   string
	NumContrato string
	CodProd     string
	CodSubProd  string
	Moneda      string
	Saldo       string
}

func SetHandlers() {

	http.HandleFunc("/cargatablaofertas", func(w http.ResponseWriter, req *http.Request) {
		err := services.CargaTablaOfertas()
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write([]byte("Tabla OFERTAS cargada exitosamente."))
		if err != nil {
			log.Fatal(err)
		}
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
			_, err = w.Write(entradaJson)
			if err != nil {
				log.Fatal(err)
			}

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
			_, err = w.Write([]byte("Insertado correctamente."))
			if err != nil {
				log.Fatal(err)
			}

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
			_, err = w.Write([]byte("Actualizado correctamente."))
			if err != nil {
				log.Fatal(err)
			}

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
			_, err = w.Write([]byte("Eliminado correctamente."))
			if err != nil {
				log.Fatal(err)
			}
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
		_, err = w.Write(entradaJson)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/contrato/{id}", func(w http.ResponseWriter, req *http.Request) {

		switch req.Method {

		case "GET":
			id := req.URL.Query().Get("id")
			body, err := services.GetContratos(id)
			if err != nil {
				log.Fatal(err)
			}
			_, err = w.Write(body)
			if err != nil {
				log.Fatal(err)
			}
		case "POST":

			var contrato Contrato
			err := json.NewDecoder(req.Body).Decode(&contrato)
			if err != nil {
				log.Fatal(err)
			}
			data := url.Values{}
			data.Add("nup", contrato.Nup)
			data.Add("codEntidad", contrato.CodEntidad)
			data.Add("codCentro", contrato.CodCentro)
			data.Add("numContrato", contrato.NumContrato)
			data.Add("codProd", contrato.CodProd)
			data.Add("codSubProd", contrato.CodSubProd)
			data.Add("moneda", contrato.Moneda)
			data.Add("saldo", contrato.Saldo)
			err = services.PostContrato(data)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
}
