package internal

import (
	"fmt"
	"github.com/Nkhinatolla/vrpSimple/vrpSimple"
	"time"
)

func main() {
	point_1 := vrpSimple.Point{
		ID:                "courier_1",
		Dependencies:      []string{},
		ShouldBeArrivedAt: time.Now().Add(3 * time.Hour),
		Latitude:          43.189297,
		Longitude:         76.871927,
	}
	point_2 := vrpSimple.Point{
		ID:                "point_1",
		Dependencies:      []string{"courier_1"},
		ShouldBeArrivedAt: time.Now().Add(3 * time.Hour),
		Latitude:          43.269373,
		Longitude:         76.936449,
	}
	point_3 := vrpSimple.Point{
		ID:                "point_2",
		Dependencies:      []string{"courier_1"},
		ShouldBeArrivedAt: time.Now().Add(3 * time.Hour),
		Latitude:          43.199297,
		Longitude:         76.871927,
	}

	points := make([]vrpSimple.Point, 3)
	points[0] = point_1
	points[1] = point_2
	points[2] = point_3

	//durations := getManualDurations()

	durations, err := getOSRMDurations(points)

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

func getOSRMDurations(points []vrpSimple.Point) ([][]int, error) {
	result, err := vrpSimple.NewOSRMService("", "walking").GetTable(points)
	if err != nil {
		return nil, err
	}

	durationsRaw, _ := result["durations"].([]interface{})

	durations := make([][]int, len(durationsRaw))
	for i, row := range durationsRaw {
		rowSlice, _ := row.([]interface{})
		durations[i] = make([]int, len(rowSlice))
		for j, val := range rowSlice {
			num, _ := val.(float64)
			durations[i][j] = int(num) / 60
		}
	}

	return durations, nil
}
