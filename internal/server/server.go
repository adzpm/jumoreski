package server

import (
	"fmt"
	"github.com/adzpm/jumoreski/models"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"gorm.io/gorm"
	"time"
)

type (
	Server struct {
		cfg     *Config
		srv     *fiber.App
		storage *gorm.DB
	}
)

func NewServer(cfg *Config, storage *gorm.DB) *Server {
	return &Server{
		cfg:     cfg,
		storage: storage,
	}
}

func (s *Server) Start() error {
	s.srv = fiber.New(fiber.Config{
		Immutable:    true,
		WriteTimeout: time.Minute,
		ReadTimeout:  time.Minute,
		JSONDecoder:  jsoniter.Unmarshal,
		JSONEncoder:  jsoniter.Marshal,
	})

	s.srv.Static("/", s.cfg.FrontPath)
	s.srv.Get("/random", s.randomHandler)

	return s.srv.Listen(fmt.Sprintf("%v:%v", s.cfg.Host, s.cfg.Port))
}

func (s *Server) randomHandler(c *fiber.Ctx) error {
	post := &models.Post{}

	if err := s.storage.Raw("SELECT * FROM posts ORDER BY RANDOM() LIMIT 1;").Scan(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("ошыбка")
	}

	return c.Status(fiber.StatusOK).JSON(post)
}
