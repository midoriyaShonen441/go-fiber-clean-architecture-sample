package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
	Fiber Fiber
}

type Fiber struct {
	Host string
	Port string
}

func LoadConfig() (*Configs, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Configs{}

	fiberPort := os.Getenv("FIBER_PORT")
	if fiberPort == "" {
		return nil, fmt.Errorf("FIBER_PORT")
	} else {
		config.Fiber.Port = fiberPort
	}

	fiberHost := os.Getenv("FIBER_HOST")
	if fiberHost == "" {
		return nil, fmt.Errorf("FIBER_HOST")
	} else {
		config.Fiber.Host = fiberHost
	}

	return &config, nil
}
