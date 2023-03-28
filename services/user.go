package services

import (
	"github.com/realtemirov/api-crud-template/models"
	"github.com/realtemirov/api-crud-template/storage"
)

type userService struct {
	db storage.UserI
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

func newUserService(repository storage.UserI) *userService {
	return &userService{
		db: repository,
	}
}

func (u *userService) CreateUser(use *models.User) (*models.User, error) {
	return u.db.CreateUser(use)
}

func (u *userService) GetUserByID(id int64) (*models.User, error) {
	return u.db.GetUserByID(id)
}

func (u *userService) GetUserByEmail(email string) (*models.User, error) {
	return u.db.GetUserByEmail(email)
}

func (u *userService) GetUserByUserName(userName string) (*models.User, error) {
	return u.db.GetUserByUserName(userName)
}

func (u *userService) GetUsers(meta *models.Meta) (*models.GetAllUsersResponse, error) {
	return u.db.GetUsers(meta)
}

func (u *userService) UpdateUser(use *models.User) (*models.User, error) {
	return u.db.UpdateUser(use)
}
func (u *userService) DeleteUser(id int64) error {
	return u.db.DeleteUser(id)
}
