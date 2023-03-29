package postgres

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/realtemirov/api-crud-template/models"
)

type userRepo struct {
	db  *sqlx.DB
	log *log.Logger
}

const (
	fieldsOfUser = "id, first_name, last_name, user_name, email, password, created_at, updated_at, deleted_at"
)

// CreateUser implements storage.UserI
func (u *userRepo) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {

	// response result
	var result models.User

	// query
	query := `
		INSERT INTO
			users (` + fieldsOfUser + `)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING ` + fieldsOfUser

	// execute query and scan result
	err := u.db.QueryRowContext(ctx, query,
		user.ID,
		user.FirstName,
		user.LastName,
		user.UserName,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
	).Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.UserName,
		&result.Email,
		&result.Password,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	// check error
	if err != nil {

		// log error
		u.log.Printf("Method CreateUser: error while creating user: %v", err)

		// return error
		return nil, err
	}

	// return result
	return &result, nil
}

// DeleteUser implements storage.UserI
func (u *userRepo) DeleteUser(ctx context.Context, id int64) (*models.User, error) {

	// response result
	var result models.User

	// query
	query := `
		UPDATE
			users
		SET
			deleted_at = $1
		WHERE
			id = $2
		RETURNING ` + fieldsOfUser

	// execute query and scan result
	err := u.db.QueryRowContext(ctx, query,
		time.Now(),
		id,
	).Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.UserName,
		&result.Email,
		&result.Password,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	// check error
	if err != nil {

		// log error
		u.log.Printf("Method DeleteUser: error while deleting user: %v", err)

		// return error
		return nil, err
	}

	// return result
	return &result, nil
}

// GetUserByEmail implements storage.UserI
func (u *userRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {

	// response result
	var result models.User

	// query
	query := `
		SELECT
			` + fieldsOfUser + `
		FROM
			users
		WHERE
			email = $1
		LIMIT 1`

	// execute query and scan result
	err := u.db.QueryRowContext(ctx, query, email).Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.UserName,
		&result.Email,
		&result.Password,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	// check error
	if err != nil {

		// log error
		u.log.Printf("Method GetUserByEmail: error while getting user by email: %v", err)

		// return error
		return nil, err
	}

	// return result
	return &result, nil
}

// GetUserByID implements storage.UserI
func (u *userRepo) GetUserByID(ctx context.Context, id int64) (*models.User, error) {

	// response result
	var result models.User

	// query
	query := `
		SELECT
			` + fieldsOfUser + `
		FROM users
		WHERE
			id = $1
		LIMIT 1`

	// execute query and scan result
	err := u.db.QueryRowContext(ctx, query, id).Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.UserName,
		&result.Email,
		&result.Password,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	// check error
	if err != nil {

		// log error
		u.log.Printf("Method GetUserByID: error while getting user by id: %v", err)

		// return error
		return nil, err
	}

	// return result
	return &result, nil
}

// GetUserByUserName implements storage.UserI
func (u *userRepo) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {

	// response result
	var result models.User

	// query
	query := `
		SELECT
			` + fieldsOfUser + `
		FROM
			users
		WHERE
			user_name = $1
		LIMIT 1`

	// execute query and scan result
	err := u.db.QueryRowContext(ctx, query,
		userName,
	).Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.UserName,
		&result.Email,
		&result.Password,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	// check error
	if err != nil {

		// log error
		u.log.Printf("Method GetUserByUserName: error while getting user by user name: %v", err)

		// return error
		return nil, err
	}

	// return result
	return &result, nil
}

// GetUsers implements storage.UserI
func (u *userRepo) GetUsers(ctx context.Context, meta *models.Meta) (*models.GetAllUsersResponse, error) {

	// response result
	var result models.GetAllUsersResponse

	// query
	query := `
		SELECT
			` + fieldsOfUser + `
		FROM
			users
		WHERE
			deleted_at IS NULL
		ORDER BY
			id DESC
		LIMIT $1
		OFFSET $2`

	// calculate limit and offset
	limit, offset := meta.GetLimitAndOffset()

	// execute query and scan result
	err := u.db.SelectContext(ctx, &result.Data, query, limit, offset)

	// check error
	if err != nil {

		// log error
		u.log.Printf("Method GetUsers: error while getting users: %v", err)

		// return error
		return nil, err
	}

	// calculate total
	query = `
		SELECT
			COUNT(id)
		FROM
			users
		WHERE
			deleted_at IS NULL`

	// execute query and scan result
	err = u.db.GetContext(ctx, &result.Meta.TotalData, query)

	// check error
	if err != nil {

		// log error
		u.log.Printf("Method GetUsers: error while getting total users: %v", err)

		// return error
		return nil, err
	}

	// calculate meta
	result.Meta = meta.SetTotalData(result.Meta.TotalData)

	// return result
	return &result, nil

}

// UpdateUser implements storage.UserI
func (u *userRepo) UpdateUser(ctx context.Context, use *models.User) (*models.User, error) {

	// response result
	var result models.User

	// query
	query := `
		UPDATE
			users
		SET
			first_name = $1,
			last_name = $2,
			user_name = $3,
			email = $4,
			password = $5,
			updated_at = $6
		WHERE
			id = $7
		RETURNING ` + fieldsOfUser

	// execute query and scan result
	err := u.db.QueryRowContext(ctx, query,
		use.FirstName,
		use.LastName,
		use.UserName,
		use.Email,
		use.Password,
		time.Now(),
		use.ID,
	).Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.UserName,
		&result.Email,
		&result.Password,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	// check error
	if err != nil {

		// log error
		u.log.Printf("Method UpdateUser: error while updating user: %v", err)

		// return error
		return nil, err
	}

	// return result
	return &result, nil
}

// newUserRepo constructor
func newUserRepo(db *sqlx.DB, log *log.Logger) *userRepo {
	return &userRepo{
		db:  db,
		log: log,
	}
}
