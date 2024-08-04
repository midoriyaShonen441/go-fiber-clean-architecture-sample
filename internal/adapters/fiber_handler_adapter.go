package adapters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midoriyaShonen441/go-fiber-clean-architecture-sample/internal/modules/samples/controllers"
	"github.com/midoriyaShonen441/go-fiber-clean-architecture-sample/internal/modules/samples/repositories"
	"github.com/midoriyaShonen441/go-fiber-clean-architecture-sample/internal/modules/samples/usecases"
)

func (s *Server) MapHandlers() error {
	// Group a version
	v1 := s.App.Group("/v1")

	//* Users group
	sampleGroup := v1.Group("/samples")
	sampleRepo := repositories.NewSampleRepo()
	sampleUseCase := usecases.NewSampleUseCase(sampleRepo)
	controllers.NewSampleController(sampleGroup, sampleUseCase)

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil
}
