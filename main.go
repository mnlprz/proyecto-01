package main

import (
	"github/mnlprz/go/proyecto-01/database"
	"github/mnlprz/go/proyecto-01/utils"
	"log"
)

func main() {
	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = utils.CargaTable(db)
	if err != nil {
		log.Fatal(err)
	}
}
