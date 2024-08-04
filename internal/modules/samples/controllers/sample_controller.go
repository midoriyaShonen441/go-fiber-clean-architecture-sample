package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midoriyaShonen441/go-fiber-clean-architecture-sample/internal/entities"
)

type SampleController struct {
	SampleUseCase entities.SampleUseCase
}

func NewSampleController(r fiber.Router, sampleUse entities.SampleUseCase) {
	controllers := &SampleController{
		SampleUseCase: sampleUse,
	}
	r.Get(
		"/",
		controllers.GetSample,
	)
}

func (h *SampleController) GetSample(c *fiber.Ctx) error {
	res, err := h.SampleUseCase.GetSample()
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      res,
	})
}
