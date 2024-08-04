package main

import (
	"github.com/midoriyaShonen441/go-fiber-clean-architecture-sample/configs"
	"github.com/midoriyaShonen441/go-fiber-clean-architecture-sample/internal/adapters"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		panic(err.Error())
	}

	s := adapters.NewServer(cfg)
	s.Start()
}
