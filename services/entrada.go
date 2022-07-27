package services

import (
	"database/sql"
	"github/mnlprz/go/proyecto-01/database"
	"github/mnlprz/go/proyecto-01/models"
	"log"
)

const selectEntrada = "SELECT * FROM entrada WHERE id = $1"

func GetEntrada(id int64) (models.Entrada, error) {

	var entrada models.Entrada

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
