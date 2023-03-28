package services

import (
	"github.com/realtemirov/api-crud-template/models"
	"github.com/realtemirov/api-crud-template/storage"
)

type postService struct {
	db storage.PostI
}

// type PostI interface {
// 	CreatePost(post *models.Post) (*models.Post, error)
// 	GetPostByID(id int64) (*models.Post, error)
// 	GetPostByUserID(userID int64) (*models.Post, error)
// 	GetPosts(meta *models.Meta) (*models.GetAllPostsResponse, error)
// 	UpdatePost(post *models.Post) (*models.Post, error)
// 	DeletePost(id int64) error
// }

func newPostService(repository storage.PostI) *postService {
	return &postService{
		db: repository,
	}
}

func (p *postService) CreatePost(post *models.Post) (*models.Post, error) {
	return p.db.CreatePost(post)
}

func (p *postService) GetPostByID(id int64) (*models.Post, error) {
	return p.db.GetPostByID(id)
}

func (p *postService) GetPostByUserID(userID int64) (*models.Post, error) {
	return p.db.GetPostByUserID(userID)
}

func (p *postService) GetPosts(meta *models.Meta) (*models.GetAllPostsResponse, error) {
	return p.db.GetPosts(meta)
}

func (p *postService) UpdatePost(post *models.Post) (*models.Post, error) {
	return p.db.UpdatePost(post)
}

func (p *postService) DeletePost(id int64) error {
	return p.db.DeletePost(id)
}
