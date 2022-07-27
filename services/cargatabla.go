package services

import (
	"encoding/csv"
	"github/mnlprz/go/proyecto-01/database"
	"io"
	"log"
	"os"
)

type Entrada struct {
	campo1 string
	campo2 string
	campo3 string
}

const createAccount = ` INSERT INTO entrada (
  campo1,
  campo2,
  campo3
) VALUES (
  $1, $2, $3
)`

func CargaTabla() error {

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	f, err := os.Open("entrada.csv")
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(f)
	csvReader.Comma = (';')
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		entrada := Entrada{
			campo1: rec[0],
			campo2: rec[1],
			campo3: rec[2],
		}
		_, err = db.Exec(createAccount, entrada.campo1, entrada.campo2, entrada.campo3)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Tabla Cargada Exitosamente...")
	return nil
}
