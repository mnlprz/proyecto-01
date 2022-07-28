package services

import (
	"database/sql"
	"github/mnlprz/go/proyecto-01/database"
	"github/mnlprz/go/proyecto-01/models"
	"log"
)

func GetEntrada(id int64) (models.Entrada, error) {

	var entrada models.Entrada
	const selectEntrada = "SELECT * FROM entrada WHERE id = $1"

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r := db.QueryRow(selectEntrada, id).Scan(&entrada.Id, &entrada.Campo1, &entrada.Campo2, &entrada.Campo3, &entrada.CreatedAt, &entrada.UpdatedAt)
	if r == sql.ErrNoRows {
		log.Fatal(r)
	}
	return entrada, nil
}

func GetEntradas(campo1 string) ([]models.Entrada, error) {

	var entradas []models.Entrada
	const selectEntrada = "SELECT * FROM entrada WHERE campo1 = $1"

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r, err := db.Query(selectEntrada, campo1)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for r.Next() {
		var entrada models.Entrada
		err = r.Scan(&entrada.Id, &entrada.Campo1, &entrada.Campo2, &entrada.Campo3, &entrada.CreatedAt, &entrada.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		entradas = append(entradas, entrada)
	}
	if r.Err() != nil {
		log.Fatal(err)
	}
	return entradas, nil
}
