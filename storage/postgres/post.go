package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/realtemirov/api-crud-template/models"
)

type postRepo struct {
	db *sqlx.DB
}

// CreatePost implements storage.PostI
func (*postRepo) CreatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	panic("unimplemented")
}

// DeletePost implements storage.PostI
func (*postRepo) DeletePost(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetPostByID implements storage.PostI
func (*postRepo) GetPostByID(ctx context.Context, id int64) (*models.Post, error) {
	panic("unimplemented")
}

// GetPostByUserID implements storage.PostI
func (*postRepo) GetPostByUserID(ctx context.Context, userID int64) (*models.Post, error) {
	panic("unimplemented")
}

// GetPosts implements storage.PostI
func (*postRepo) GetPosts(ctx context.Context, meta *models.Meta) (*models.GetAllPostsResponse, error) {
	panic("unimplemented")
}

// UpdatePost implements storage.PostI
func (*postRepo) UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	panic("unimplemented")
}

// newPostRepo constructor
func newPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}
