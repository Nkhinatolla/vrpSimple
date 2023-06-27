package vrpSimple

import (
	"encoding/json"
	"fmt"
	"io"
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

func (s *OSRMService) PointsToFormattedString(points []Point) string {
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

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

func (s *OSRMService) GetRoutes(points []Point) (map[string]interface{}, error) {
	pointsStr := s.PointsToFormattedString(points)
	urlEnd := s.RoutePath + s.Profile + pointsStr
	return s.SendRequest(urlEnd)
}

func (s *OSRMService) GetTrips(points []Point) (map[string]interface{}, error) {
	pointsStr := s.PointsToFormattedString(points)
	urlEnd := s.TripPath + s.Profile + pointsStr + "?source=first"
	return s.SendRequest(urlEnd)
}

func (s *OSRMService) GetTable(points []Point) (map[string]interface{}, error) {
	pointsStr := s.PointsToFormattedString(points)
	urlEnd := s.TablePath + s.Profile + pointsStr
	return s.SendRequest(urlEnd)
}

func (s *OSRMService) GetDurationsByTable(points []Point) ([][]int, error) {
	result, err := s.GetTable(points)
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
