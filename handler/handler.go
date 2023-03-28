package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/realtemirov/api-crud-template/config"
	"github.com/realtemirov/api-crud-template/services"
)

type Handler struct {
	log      *log.Logger
	cnf      *config.Config
	services *services.Services
}

func NewHandler(cnf *config.Config, services *services.Services, log *log.Logger) *Handler {
	return &Handler{
		log:      log,
		cnf:      cnf,
		services: services,
	}
}
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
