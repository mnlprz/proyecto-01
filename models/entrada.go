package models

import "time"

type Entrada struct {
	Id        int64
	Campo1    string
	Campo2    string
	Campo3    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
