package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/realtemirov/api-crud-template/models"
)

type userRepo struct {
	db *sqlx.DB
}

// CreateUser implements storage.UserI
func (*userRepo) CreateUser(ctx context.Context, use *models.User) (*models.User, error) {
	panic("unimplemented")
}

// DeleteUser implements storage.UserI
func (*userRepo) DeleteUser(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetUserByEmail implements storage.UserI
func (*userRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	panic("unimplemented")
}

// GetUserByID implements storage.UserI
func (*userRepo) GetUserByID(ctx context.Context, id int64) (*models.User, error) {
	panic("unimplemented")
}

// GetUserByUserName implements storage.UserI
func (*userRepo) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {
	panic("unimplemented")
}

// GetUsers implements storage.UserI
func (*userRepo) GetUsers(ctx context.Context, meta *models.Meta) (*models.GetAllUsersResponse, error) {
	panic("unimplemented")
}

// UpdateUser implements storage.UserI
func (*userRepo) UpdateUser(ctx context.Context, use *models.User) (*models.User, error) {
	panic("unimplemented")
}

// newUserRepo constructor
func newUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}
