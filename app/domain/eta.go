package domain

import "time"

type EtaPoint struct {
	ID                string
	Dependencies      []string
	ShouldBeArrivedAt time.Time
	EstimateAt        time.Time
	Priority          int
	Latitude          float64
	Longitude         float64
}
