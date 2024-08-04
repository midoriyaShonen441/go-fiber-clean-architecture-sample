package repositories

import "github.com/midoriyaShonen441/go-fiber-clean-architecture-sample/internal/entities"

type SampleRepo struct {
}

func NewSampleRepo() entities.SampleRepo {
	return &SampleRepo{}
}

func (s *SampleRepo) GetSample() (*entities.SampleRes, error) {
	return &entities.SampleRes{
		Data: entities.Sample{
			Name: "Natta",
			Age:  28,
		}}, nil
}
