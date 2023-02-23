package service

type HealthService struct {
}

type IHealthService interface {
	HealthCheck() error
}

func NewHealthService() IHealthService {
	return &HealthService{
		//db: repo,
	}
}

func (h *HealthService) HealthCheck() error {
	return nil
}
