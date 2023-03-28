package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/realtemirov/api-crud-template/config"
	"github.com/realtemirov/api-crud-template/handler"
)

func main() {
	log.Println("Starting the application...")

	cnf := config.NewConfig()

	r := gin.Default()

	h := handler.NewHandler()
	r.GET("/ping", h.Ping)

	r.Run(fmt.Sprintf(":%s", cnf.HostPort))
}
