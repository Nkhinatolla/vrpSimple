package services

import (
	"git.chocofood.kz/chocodelivery/assignments/vrp-simple/domain"
)

type EtaService struct {
	Points               []domain.EtaPoint
	Host                 string
	Profile              string
	PickupLagTime        int
	TravelTimeMultiplier float64
}

func NewEtaService(points []domain.EtaPoint, host, profile string, pickupLagTime int, travelTimeMultiplier float64) *EtaService {
	return &EtaService{
		Points:               points,
		Host:                 host,
		Profile:              profile,
		PickupLagTime:        pickupLagTime,
		TravelTimeMultiplier: travelTimeMultiplier,
	}
}

func (s *EtaService) FindOptimalEta(ignoreShouldArrivedAt bool) ([]domain.EtaPoint, error) {
	durations, err := s.GetDurations()
	if err != nil {
		return nil, err
	}

	algorithmService := NewEtaAlgorithmService(s.Points)
	return algorithmService.Calculate(durations, false)
}

func (s *EtaService) GetDurations() ([][]int, error) {
	osrmService := NewOSRMService(s.Host, s.Profile)
	durations, err := osrmService.GetDurationsByTable(s.Points)
	for i := range durations {
		for j := range durations[i] {
			durations[i][j] = int(float64(durations[i][j])*s.TravelTimeMultiplier) + s.PickupLagTime
		}
	}
	return durations, err
}
