package vrp_simple

import (
	"fmt"
	"time"
)

func main() {
	// Comment
	point_1 := Point{
		ID:                "courier_1",
		Dependencies:      []string{},
		ShouldBeArrivedAt: time.Now().Add(3 * time.Hour),
		Latitude:          43.189297,
		Longitude:         76.871927,
	}
	point_2 := Point{
		ID:                "point_1",
		Dependencies:      []string{"courier_1"},
		ShouldBeArrivedAt: time.Now().Add(3 * time.Hour),
		Latitude:          43.269373,
		Longitude:         76.936449,
	}
	point_3 := Point{
		ID:                "point_2",
		Dependencies:      []string{"courier_1"},
		ShouldBeArrivedAt: time.Now().Add(3 * time.Hour),
		Latitude:          43.199297,
		Longitude:         76.871927,
	}

	points := make([]Point, 3)
	points[0] = point_1
	points[1] = point_2
	points[2] = point_3

	//durations := getManualDurations()
	osrmService := NewOSRMService("", "walking")
	durations, err := osrmService.GetDurationsByTable(points)

	fmt.Println(durations)

	if err != nil {
		fmt.Println(err)
	}

	points, err = NewEtaAlgorithmService(points).Calculate(durations, false)
	if err != nil {
		fmt.Println(err)
	}
	for _, point := range points {
		fmt.Println(point.Priority, point.ID, point.EstimateAt)
	}
}

func getManualDurations() [][]int {
	durations := make([][]int, 3)
	durations[0] = append(durations[0], []int{0, 141, 10}...)
	durations[1] = append(durations[1], []int{141, 0, 140}...)
	durations[2] = append(durations[2], []int{10, 140, 0}...)
	return durations
}
