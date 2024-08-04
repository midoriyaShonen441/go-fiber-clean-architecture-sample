package utils

import (
	"errors"
	"fmt"

	"github.com/midoriyaShonen441/go-fiber-clean-architecture-sample/configs"
)

func ConnectionUrlBuilder(stuff string, cfg *configs.Configs) (string, error) {
	var url string

	switch stuff {
	case "fiber":
		url = fmt.Sprintf("%s:%s", cfg.Fiber.Host, cfg.Fiber.Port)
	default:
		errMsg := fmt.Sprintf("error, connection url builder doesn't know the %s", stuff)
		return "", errors.New(errMsg)
	}
	return url, nil
}
