package entities

type Sample struct {
	Name string
	Age  int
}

type SampleRes struct {
	Data Sample `json:"data"`
}

type SampleUseCase interface {
	GetSample() (*SampleRes, error)
}

type SampleRepo interface {
	GetSample() (*SampleRes, error)
}
