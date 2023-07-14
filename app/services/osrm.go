package services

import (
	"encoding/json"
	"fmt"
	"git.chocofood.kz/chocodelivery/assignments/vrp-simple/domain"
	"net/http"
	"strings"
)

type OSRMService struct {
	RoutePath              string
	TripPath               string
	TablePath              string
	SourceQuery            string
	DestinationQuery       string
	DistanceQuery          string
	SuccessfulResponseCode string
	Host                   string
	Profile                string
}

func NewOSRMService(host string, profile string) *OSRMService {

	return &OSRMService{
		RoutePath:              "/route/v1/",
		TripPath:               "/trip/v1/",
		TablePath:              "/table/v1/",
		SourceQuery:            "sources=",
		DestinationQuery:       "destinations=",
		DistanceQuery:          "annotations=distance",
		SuccessfulResponseCode: "Ok",
		Host:                   host,
		Profile:                profile + "/",
	}
}

func (s *OSRMService) PointsToFormattedString(points []domain.EtaPoint) string {
	formattedPoints := make([]string, len(points))
	for i, point := range points {
		formattedPoints[i] = fmt.Sprintf("%.6f,%.6f", point.Longitude, point.Latitude)
	}
	formattedString := strings.Join(formattedPoints, ";")
	return formattedString
}

func (s *OSRMService) SendRequest(urlEnd string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s%s", s.Host, urlEnd)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	code, ok := result["code"].(string)
	if !ok || code != s.SuccessfulResponseCode {
		message, _ := result["message"].(string)
		return nil, fmt.Errorf("unprocessable Entity: %s", message)
	}

	return result, nil
}

func (s *OSRMService) GetRoutes(points []domain.EtaPoint) (domain.OSRMRoute, error) {
	pointsStr := s.PointsToFormattedString(points)
	urlEnd := s.RoutePath + s.Profile + pointsStr

	result, err := s.SendRequest(urlEnd)
	responseBytes, _ := json.Marshal(result)
	response := domain.OSRMRoute{}
	err = json.Unmarshal(responseBytes, &response)

	if err != nil {
		return domain.OSRMRoute{}, err
	}
	return response, nil
}

func (s *OSRMService) GetTrips(points []domain.EtaPoint) (domain.OSRMTrip, error) {
	pointsStr := s.PointsToFormattedString(points)
	urlEnd := s.TripPath + s.Profile + pointsStr + "?source=first"

	result, err := s.SendRequest(urlEnd)
	responseBytes, _ := json.Marshal(result)

	response := domain.OSRMTrip{}
	err = json.Unmarshal(responseBytes, &response)

	if err != nil {
		return domain.OSRMTrip{}, err
	}
	return response, nil
}

func (s *OSRMService) GetTable(points []domain.EtaPoint) (domain.OSRMTable, error) {
	pointsStr := s.PointsToFormattedString(points)
	urlEnd := s.TablePath + s.Profile + pointsStr

	result, err := s.SendRequest(urlEnd)
	responseBytes, _ := json.Marshal(result)
	response := domain.OSRMTable{}
	err = json.Unmarshal(responseBytes, &response)

	if err != nil {
		return domain.OSRMTable{}, err
	}
	return response, nil
}

func (s *OSRMService) GetDurationsByTable(points []domain.EtaPoint) ([][]int, error) {
	result, err := s.GetTable(points)
	if err != nil {
		return nil, err
	}
	durations := ConvertFloat64ToInt(result.Durations)
	for i := range durations {
		for j := range durations[i] {
			durations[i][j] = durations[i][j] / 60
		}
	}

	return durations, nil
}

func ConvertFloat64ToInt(slice [][]float64) [][]int {
	result := make([][]int, len(slice))
	for i := range slice {
		result[i] = make([]int, len(slice[i]))
		for j := range slice[i] {
			result[i][j] = int(slice[i][j])
		}
	}
	return result
}
