package internal

import (
	"fmt"
	"github.com/Nkhinatolla/vrpSimple/vrpSimple"
	"time"
)

type EtaAlgorithmService struct {
	InfDatetime time.Time
	PointIndex  map[string]int
	Points      []vrpSimple.Point
}

func reverse(arr []vrpSimple.Point) {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		j := length - 1 - i
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func NewEtaAlgorithmService(points []vrpSimple.Point) *EtaAlgorithmService {
	service := &EtaAlgorithmService{
		InfDatetime: time.Date(2100, 10, 14, 0, 0, 0, 0, time.UTC),
		PointIndex:  make(map[string]int),
		Points:      make([]vrpSimple.Point, 0),
	}

	for _, point := range points {
		service.AddPoint(point)
	}

	return service
}

func (s *EtaAlgorithmService) AddPoint(point vrpSimple.Point) {
	s.Points = append(s.Points, point)
	s.PointIndex[point.ID] = len(s.Points) - 1
}

func (s *EtaAlgorithmService) GetPoint(id string) *int {
	if idx, ok := s.PointIndex[id]; ok {
		return &idx
	}
	return nil
}

func (s *EtaAlgorithmService) DependenciesCheck(mask int, masks []int, point vrpSimple.Point) bool {
	for _, id := range point.Dependencies {
		if pointID := s.GetPoint(id); pointID != nil {
			if mask&masks[*pointID] == 0 {
				return false
			}
		}
	}
	return true
}

func (s *EtaAlgorithmService) Calculate(durations [][]int, ignoreShouldArrivedAt bool) ([]vrpSimple.Point, error) {
	nodeCount := len(s.Points)
	now := time.Now().UTC()
	masks := make([]int, nodeCount)
	for i := 0; i < nodeCount; i++ {
		masks[i] = 1 << i
	}
	allVisitedMask := (1 << nodeCount) - 1

	queue := make([][]int, 0)
	for i := 0; i < nodeCount; i++ {
		if len(s.Points[i].Dependencies) == 0 {
			queue = append(queue, []int{i, masks[i]})
		}
	}

	dp := make(map[[2]int][2]interface{})
	for _, item := range queue {
		currentNode := item[0]
		visitedMask := item[1]
		dp[[2]int{visitedMask, currentNode}] = [2]interface{}{now, currentNode}
	}

	for len(queue) > 0 {
		currentItem := queue[0]
		queue = queue[1:]
		currentNode := currentItem[0]
		visitedMask := currentItem[1]

		if visitedMask == allVisitedMask {
			continue
		}
		for neighbor := 0; neighbor < nodeCount; neighbor++ {
			newVisitMask := visitedMask | masks[neighbor]

			if newVisitMask != visitedMask && s.DependenciesCheck(newVisitMask, masks, s.Points[neighbor]) {
				duration := time.Duration(durations[currentNode][neighbor]) * time.Minute
				currentDatetime := dp[[2]int{visitedMask, currentNode}][0].(time.Time).Add(duration)

				to := [2]int{newVisitMask, neighbor}

				oldData, exists := dp[to]
				if !exists {
					oldData = [2]interface{}{s.InfDatetime, nil}
				}
				oldDatetime := oldData[0].(time.Time)

				shouldArrived := s.Points[neighbor].ShouldBeArrivedAt
				if ignoreShouldArrivedAt {
					shouldArrived = s.InfDatetime
				}
				if currentDatetime.Before(oldDatetime) && currentDatetime.Before(shouldArrived) {
					dp[to] = [2]interface{}{currentDatetime, currentNode}
					queue = append(queue, []int{neighbor, newVisitMask})
				}
			}
		}
	}

	answer := [2]int{-1, -1}
	for i := 0; i < nodeCount; i++ {
		to := [2]int{allVisitedMask, i}
		if val, ok := dp[to]; ok && (answer[0] == -1 || dp[answer][0].(time.Time).After(val[0].(time.Time))) {
			answer = to
		}
	}

	if answer[0] == -1 {
		return nil, fmt.Errorf("there is no route ")
	}

	result := make([]vrpSimple.Point, 0)
	current := dp[answer]
	lastNode := answer[1]
	for current != [2]interface{}{} {
		s.Points[lastNode].EstimateAt = current[0].(time.Time)
		result = append(result, s.Points[lastNode])
		allVisitedMask ^= masks[lastNode]
		lastNode = current[1].(int)
		current = dp[[2]int{allVisitedMask, lastNode}]
	}

	reverse(result)

	for i := range result {
		result[i].Priority = i
	}

	return result, nil
}
