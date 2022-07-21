package models

import "time"

type SolicitudRefi struct {
	Id        int64
	Monto     float64
	Estado    string
	Contratos []Contrato
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (solicitudRefi *SolicitudRefi) AltaSolicitudRefi() error {

	return nil
}
