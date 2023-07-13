package domain

type ShortPoint struct {
	Distance float64   `json:"distance"`
	Hint     string    `json:"hint"`
	Location []float64 `json:"location"`
	Name     string    `json:"name"`
}

type OSRMRoute struct {
	Code   string `json:"code"`
	Routes []struct {
		Distance float64 `json:"distance"`
		Duration float64 `json:"duration"`
		Geometry string  `json:"geometry"`
		Legs     []struct {
			Distance float64       `json:"distance"`
			Duration float64       `json:"duration"`
			Steps    []interface{} `json:"steps"`
			Summary  string        `json:"summary"`
			Weight   float64       `json:"weight"`
		} `json:"legs"`
		Weight     float64 `json:"weight"`
		WeightName string  `json:"weight_name"`
	} `json:"routes"`
	Waypoints []ShortPoint `json:"waypoints"`
}

type OSRMTrip struct {
	Code  string `json:"code"`
	Trips []struct {
		Distance float64 `json:"distance"`
		Duration float64 `json:"duration"`
		Geometry string  `json:"geometry"`
		Legs     []struct {
			Distance float64       `json:"distance"`
			Duration float64       `json:"duration"`
			Steps    []interface{} `json:"steps"`
			Summary  string        `json:"summary"`
			Weight   float64       `json:"weight"`
		} `json:"legs"`
		Weight     float64 `json:"weight"`
		WeightName string  `json:"weight_name"`
	} `json:"routes"`
	Waypoints struct {
		WaypointIndex int       `json:"waypoint_index"`
		TripIndex     int       `json:"trips_index"`
		Distance      float64   `json:"distance"`
		Hint          string    `json:"hint"`
		Location      []float64 `json:"location"`
		Name          string    `json:"name"`
	} `json:"waypoints"`
}

type OSRMTable struct {
	Code         string       `json:"code"`
	Destinations []ShortPoint `json:"routes"`
	Durations    [][]float64
	Sources      []ShortPoint `json:"waypoints"`
}
