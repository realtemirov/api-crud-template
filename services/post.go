package services

import (
	"log"

	"github.com/realtemirov/api-crud-template/storage"
)

type postService struct {
	log  *log.Logger
	repo storage.PostI
}

func newPostService(repo storage.PostI, log *log.Logger) *postService {
	return &postService{
		log:  log,
		repo: repo,
	}
}
