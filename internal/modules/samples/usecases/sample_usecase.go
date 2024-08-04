package usecases

import "github.com/midoriyaShonen441/go-fiber-clean-architecture-sample/internal/entities"

type SampleUseCase struct {
	Repo entities.SampleRepo
}

func NewSampleUseCase(repo entities.SampleRepo) entities.SampleUseCase {
	return &SampleUseCase{Repo: repo}
}

func (s *SampleUseCase) GetSample() (*entities.SampleRes, error) {
	return s.GetSample()
}
