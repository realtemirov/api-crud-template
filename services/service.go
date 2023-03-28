package services

import "github.com/realtemirov/api-crud-template/storage"

type service struct {
	db storage.StorageI

	UserService *userService
	PostService *postService
}

func NewService(repository storage.StorageI) *service {
	return &service{
		db:          repository,
		UserService: newUserService(repository.User()),
		PostService: newPostService(repository.Post()),
	}
}
