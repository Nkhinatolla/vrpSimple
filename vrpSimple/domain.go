package vrpSimple

import "time"

type Point struct {
	ID                string
	Dependencies      []string
	ShouldBeArrivedAt time.Time
	EstimateAt        time.Time
	Priority          int
	Latitude          float64
	Longitude         float64
}
