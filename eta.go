package vrp_simple

type EtaService struct {
	Points        []Point
	Host          string
	Profile       string
	PickupLagTime int
}

func (s *EtaService) FindOptimalEta() ([]Point, error) {
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
	return durations, err
}
