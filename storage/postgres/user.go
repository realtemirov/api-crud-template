package postgres

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/realtemirov/gorm/models"
)

type userRepo struct {
	db *sqlx.DB
}

// CreateUser implements storage.UserI
func (u *userRepo) CreateUser(use *models.User) (*models.User, error) {
	panic("unimplemented")
}

// DeleteUser implements storage.UserI
func (u *userRepo) DeleteUser(id int64) error {
	panic("unimplemented")
}

// GetUserByEmail implements storage.UserI
func (u *userRepo) GetUserByEmail(email string) (*models.User, error) {
	panic("unimplemented")
}

// GetUserByID implements storage.UserI
func (u *userRepo) GetUserByID(id int64) (*models.User, error) {
	panic("unimplemented")
}

// GetUserByUserName implements storage.UserI
func (u *userRepo) GetUserByUserName(userName string) (*models.User, error) {
	panic("unimplemented")
}

// GetUsers implements storage.UserI
func (u *userRepo) GetUsers(meta *models.Meta) (*models.GetAllUsersResponse, error) {
	panic("unimplemented")
}

// UpdateUser implements storage.UserI
func (u *userRepo) UpdateUser(use *models.User) (*models.User, error) {
	panic("unimplemented")
}

// NewUserRepo constructor
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}
