package models

import "time"

type Oferta struct {
	NumPersona string
	Oferta     string
	Nombre     string
	Apellido   string
	SitIva     string
	Digital    string
	Fecha      string
	Created_at time.Time
	Updated_at time.Time
}
