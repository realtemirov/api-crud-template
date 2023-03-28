package services

import (
	"log"

	"github.com/realtemirov/api-crud-template/storage"
)

type userService struct {
	log  *log.Logger
	repo storage.UserI
}

//	type UserI interface {
//		CreateUser(use *models.User) (*models.User, error)
//		GetUserByID(id int64) (*models.User, error)
//		GetUserByEmail(email string) (*models.User, error)
//		GetUserByUserName(userName string) (*models.User, error)
//		GetUsers(meta *models.Meta) (*models.GetAllUsersResponse, error)
//		UpdateUser(use *models.User) (*models.User, error)
//		DeleteUser(id int64) error
//	}

func newUserService(repo storage.UserI, log *log.Logger) *userService {
	return &userService{
		log:  log,
		repo: repo,
	}
}
