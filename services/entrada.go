package services

import (
	"database/sql"
	"github/mnlprz/go/proyecto-01/database"
	"github/mnlprz/go/proyecto-01/models"
	"log"
)

func GetEntrada(id int64) (models.Entrada, error) {

	const query = "SELECT * FROM entrada WHERE id = $1"
	var entrada models.Entrada

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := db.QueryRow(query, id).Scan(&entrada.Id, &entrada.Campo1, &entrada.Campo2, &entrada.Campo3, &entrada.CreatedAt, &entrada.UpdatedAt)
	if r == sql.ErrNoRows {
		log.Fatal(r)
	}

	return entrada, nil
}

func DeleteEntrada(id int64) error {

	const query = "DELETE FROM entrada WHERE id = $1"

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func UpdateEntrada(entrada models.Entrada) error {

	const query = "UPDATE entrada SET campo1 = $1, campo2 = $2, campo3 = $3, updated_at = NOW() WHERE id = $4"

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(query, entrada.Campo1, entrada.Campo2, entrada.Campo3, entrada.Id)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func InsertEntrada(entrada models.Entrada) error {

	const query = "INSERT INTO entrada (campo1, campo2, campo3, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())"

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(query, entrada.Campo1, entrada.Campo2, entrada.Campo3)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func GetEntradas(campo1 string) ([]models.Entrada, error) {

	const query = "SELECT * FROM entrada WHERE campo1 = $1"
	var entradas []models.Entrada

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r, err := db.Query(query, campo1)
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
