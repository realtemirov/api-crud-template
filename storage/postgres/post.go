package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/realtemirov/api-crud-template/models"
)

type postRepo struct {
	db *sqlx.DB
}

// CreatePost implements storage.PostI
func (p *postRepo) CreatePost(post *models.Post) (*models.Post, error) {
	panic("unimplemented")
}

// DeletePost implements storage.PostI
func (p *postRepo) DeletePost(id int64) error {
	panic("unimplemented")
}

// GetPostByID implements storage.PostI
func (p *postRepo) GetPostByID(id int64) (*models.Post, error) {
	panic("unimplemented")
}

// GetPostByUserID implements storage.PostI
func (p *postRepo) GetPostByUserID(userID int64) (*models.Post, error) {
	panic("unimplemented")
}

// GetPosts implements storage.PostI
func (p *postRepo) GetPosts(meta *models.Meta) (*models.GetAllPostsResponse, error) {
	panic("unimplemented")
}

// UpdatePost implements storage.PostI
func (p *postRepo) UpdatePost(post *models.Post) (*models.Post, error) {
	panic("unimplemented")
}

// NewPostRepo constructor
func NewPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}
