package storage

import "github.com/realtemirov/api-crud-template/models"

type StorageI interface {
	CloseDB() error
	User() UserI
	Post() PostI
}

type UserI interface {
	CreateUser(use *models.User) (*models.User, error)
	GetUserByID(id int64) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUserName(userName string) (*models.User, error)
	GetUsers(meta *models.Meta) (*models.GetAllUsersResponse, error)
	UpdateUser(use *models.User) (*models.User, error)
	DeleteUser(id int64) error
}

type PostI interface {
	CreatePost(post *models.Post) (*models.Post, error)
	GetPostByID(id int64) (*models.Post, error)
	GetPostByUserID(userID int64) (*models.Post, error)
	GetPosts(meta *models.Meta) (*models.GetAllPostsResponse, error)
	UpdatePost(post *models.Post) (*models.Post, error)
	DeletePost(id int64) error
}
