package adapters

import (
	"github.com/midoriyaShonen441/go-fiber-clean-architecture-sample/pkg/utils"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/midoriyaShonen441/go-fiber-clean-architecture-sample/configs"
)

type Server struct {
	App *fiber.App
	Cfg *configs.Configs
}

func NewServer(cfg *configs.Configs) *Server {
	return &Server{
		App: fiber.New(),
		Cfg: cfg,
	}
}

func (s *Server) Start() {
	if err := s.MapHandlers(); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	fiberConnURL, err := utils.ConnectionUrlBuilder("fiber", s.Cfg)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	host := s.Cfg.Fiber.Host
	port := s.Cfg.Fiber.Port
	log.Printf("server has been started on %s:%s ⚡", host, port)

	if err := s.App.Listen(fiberConnURL); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
}
