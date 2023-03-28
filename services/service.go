package services

import (
	"log"

	"github.com/realtemirov/api-crud-template/config"
	"github.com/realtemirov/api-crud-template/storage"
)

type Services struct {
	storage     storage.StorageI
	cnf         *config.Config
	log         *log.Logger
	UserService *userService
	PostService *postService
}

func NewService(cnf *config.Config, storage storage.StorageI, log *log.Logger) *Services {
	return &Services{
		storage:     storage,
		cnf:         cnf,
		log:         log,
		UserService: newUserService(storage.User(), log),
		PostService: newPostService(storage.Post(), log),
	}
}
