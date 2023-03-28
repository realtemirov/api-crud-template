package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/realtemirov/api-crud-template/config"
	"github.com/realtemirov/api-crud-template/handler"
	"github.com/realtemirov/api-crud-template/services"
	"github.com/realtemirov/api-crud-template/storage/postgres"
)

func main() {

	// init logger
	log := log.New(nil, "Init", 0)

	// config
	cnf := config.NewConfig()

	// context
	ctx := context.Background()

	// database
	db, err := postgres.NewPostgres(ctx, cnf, log)

	// checking for error
	if err != nil {
		log.Fatalf("Method: main Comment: db.NewPostgres Error: %v", err)
	}

	// close database connection
	defer db.CloseDB()

	// services
	services := services.NewService(cnf, db, log)

	// handler
	h := handler.NewHandler(cnf, services, log)

	r := gin.Default()
	r.GET("/ping", h.Ping)
	r.Run(fmt.Sprintf(":%s", cnf.HostPort))
}
