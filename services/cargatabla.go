package services

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/mnlprz/go/proyecto-01/database"
	"github.com/mnlprz/go/proyecto-01/models"
)

func CargaTablaOfertas() error {

	const QUERY_OFERTAS = "INSERT INTO ofertas (num_persona, oferta, nombre, apellido, sit_iva, digital, fecha, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())"

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	f, err := os.Open("ofertas.csv")
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(f)
	csvReader.Comma = (',')
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		oferta := models.Oferta{
			NumPersona: rec[0],
			Oferta:     rec[1],
			Nombre:     rec[2],
			Apellido:   rec[3],
			SitIva:     rec[4],
			Digital:    rec[5],
			Fecha:      rec[6],
		}
		_, err = db.Exec(QUERY_OFERTAS, oferta.NumPersona, oferta.Oferta, oferta.Nombre, oferta.Apellido, oferta.SitIva, oferta.Digital, oferta.Fecha)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Tabla O F E R T A S cargada exitosamente...")
	return nil

}
