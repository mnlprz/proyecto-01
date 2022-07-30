package models

import "time"

type Persona struct {
	Id          int64
	NumPersona  string
	CodEntidad  string
	CodCentro   string
	NumContrato string
	CodProd     string
	CodSubProd  string
	Moneda      string
	Saldo       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
