package postgres

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/realtemirov/api-crud-template/models"
)

type postRepo struct {
	db  *sqlx.DB
	log *log.Logger
}

const (
	fieldsOfPost = "id, user_id, title, body, created_at, updated_at, deleted_at"
)

// CreatePost implements storage.PostI
func (p *postRepo) CreatePost(ctx context.Context, post *models.Post) (*models.Post, error) {

	// response result
	var result models.Post

	// query
	query := `
		INSERT INTO 
			posts (` + fieldsOfPost + `)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING ` + fieldsOfPost

	// execute query and scan result
	err := p.db.QueryRowContext(ctx, query,
		post.ID,
		post.UserID,
		post.Title,
		post.Body,
		post.CreatedAt,
		post.UpdatedAt,
		post.DeletedAt,
	).Scan(
		&result.ID,
		&result.UserID,
		&result.Title,
		&result.Body,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	// check error
	if err != nil {

		// log error
		p.log.Printf("Method: CreatePost Error: %v", err)

		// return error
		return nil, err
	}

	// if no error, return result
	return &result, nil
}

// DeletePost implements storage.PostI
func (p *postRepo) DeletePost(ctx context.Context, id int64) (*models.Post, error) {

	// response result
	var result models.Post

	// query
	query := `
		UPDATE 
			posts 
		SET 
			deleted_at = $1 
		WHERE 
			id = $2 
		RETURNING ` + fieldsOfPost

	// execute query and scan result
	err := p.db.QueryRowContext(ctx, query,
		time.Now().Unix(),
		id,
	).Scan(
		&result.ID,
		&result.UserID,
		&result.Title,
		&result.Body,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	// check error
	if err != nil {

		// log error
		p.log.Printf("Method: DeletePost Error: %v", err)

		// return error
		return nil, err
	}

	// if no error, return result
	return &result, nil
}

// GetPostByID implements storage.PostI
func (p *postRepo) GetPostByID(ctx context.Context, id int64) (*models.Post, error) {

	// response result
	var result models.Post

	// query
	query := `
		SELECT
			` + fieldsOfPost + `
		FROM
			posts
		WHERE
			id = $1
		AND
			deleted_at IS NULL`

	// execute query and scan result
	err := p.db.QueryRowContext(ctx, query,
		id,
	).Scan(
		&result.ID,
		&result.UserID,
		&result.Title,
		&result.Body,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	// check error
	if err != nil {

		// log error
		p.log.Printf("Method: GetPostByID Error: %v", err)

		// return error
		return nil, err
	}

	// if no error, return result
	return &result, nil
}

// GetPostByUserID implements storage.PostI
func (p *postRepo) GetPostByUserID(ctx context.Context, userID int64) (*models.Post, error) {

	// response result
	var result models.Post

	// query
	query := `
		SELECT
			` + fieldsOfPost + `
		FROM
			posts
		WHERE
			user_id = $1
		AND
			deleted_at IS NULL`

	// execute query and scan result
	err := p.db.QueryRowContext(ctx, query,
		userID,
	).Scan(
		&result.ID,
		&result.UserID,
		&result.Title,
		&result.Body,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	// check error
	if err != nil {

		// log error
		p.log.Printf("Method: GetPostByUserID Error: %v", err)

		// return error
		return nil, err
	}

	// if no error, return result
	return &result, nil
}

// GetPosts implements storage.PostI
func (p *postRepo) GetPosts(ctx context.Context, meta *models.Meta) (*models.GetAllPostsResponse, error) {

	// response result
	var result models.GetAllPostsResponse

	// query
	query := `
		SELECT
			` + fieldsOfPost + `
		FROM
			posts
		WHERE
			deleted_at IS NULL
		ORDER BY
			id DESC
		LIMIT
			$1
		OFFSET
			$2`

	// calculate limit and offset
	limit, offset := meta.GetLimitAndOffset()

	// execute query and scan result
	err := p.db.SelectContext(ctx, &result.Data, query,
		limit,
		offset,
	)

	// check error
	if err != nil {

		// log error
		p.log.Printf("Method: GetPosts Error: %v", err)

		// return error
		return nil, err
	}

	// // close rows
	// defer rows.Close()

	// // loop rows
	// for rows.Next() {

	// 	// response result
	// 	var post models.Post

	// 	// scan result
	// 	err = rows.Scan(
	// 		&post.ID,
	// 		&post.UserID,
	// 		&post.Title,
	// 		&post.Body,
	// 		&post.CreatedAt,
	// 		&post.UpdatedAt,
	// 		&post.DeletedAt,
	// 	)

	// 	// check error
	// 	if err != nil {

	// 		// log error
	// 		p.log.Printf("Method: GetPosts Error: %v", err)

	// 		// return error
	// 		return nil, err
	// 	}

	// 	// append result
	// 	result.Data = append(result.Data, &post)
	// }

	// calculate total
	query = `
		SELECT
			COUNT(id)
		FROM
			posts
		WHERE
			deleted_at IS NULL`

	// execute query and scan result
	err = p.db.GetContext(ctx, &result.Meta.TotalData, query)

	// check error
	if err != nil {

		// log error
		p.log.Printf("Method: GetPosts Error: %v", err)

		// return error
		return nil, err
	}

	// calculate meta
	result.Meta = meta.SetTotalData(result.Meta.TotalData)

	// if no error, return result
	return &result, nil
}

// UpdatePost implements storage.PostI
func (p *postRepo) UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error) {

	// response result
	var result models.Post

	// query
	query := `
		UPDATE
			posts
		SET
			user_id = $1,
			title = $2,
			body = $3,
			updated_at = $4
		WHERE
			id = $5
		RETURNING ` + fieldsOfPost

	// execute query and scan result
	err := p.db.QueryRowContext(ctx, query,
		post.UserID,
		post.Title,
		post.Body,
		time.Now().Unix(),
		post.ID,
	).Scan(
		&result.ID,
		&result.UserID,
		&result.Title,
		&result.Body,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	// check error
	if err != nil {

		// log error
		p.log.Printf("Method: UpdatePost Error: %v", err)

		// return error
		return nil, err
	}

	// if no error, return result
	return &result, nil
}

// newPostRepo constructor
func newPostRepo(db *sqlx.DB, log *log.Logger) *postRepo {
	return &postRepo{
		db:  db,
		log: log,
	}
}
