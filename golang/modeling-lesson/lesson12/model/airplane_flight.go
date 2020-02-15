package model

import (
	"time"
)

// 航空便
type AirplaneFlight struct {
	ID                int
	Name              string    // 便名
	DepartureLocation string    // 出発地
	DepartedAt        time.Time // 出発時刻
	ArrivalLocation   string    // 到着地
	ArrivedAt         time.Time // 到着時刻
}

func NewAirplaneFlight(name, departureLocation string, departedAt time.Time, arrivalLocation string, arrivedAt time.Time) *AirplaneFlight {
	return &AirplaneFlight{
		ID:                1,
		Name:              name,
		DepartureLocation: departureLocation,
		DepartedAt:        departedAt,
		ArrivalLocation:   arrivalLocation,
		ArrivedAt:         arrivedAt,
	}
}
