package model

import "time"

type Branch struct {
	ID          int    // уникальный идентификатор филиала
	Department  string // наименование филиала
	Address     string // адрес филиала
	Contract    string
	Aria        float64   // занимаемая площадь
	MeterInYear float64   // стоимость м2 в год
	TotalInYear float64   // общая стоимость в год
	UpdatedAt   time.Time // дата последнего обновления или создания
}
