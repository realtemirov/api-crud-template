package models

type Base struct {
	ID        int64  `json:"id" db:"id"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
	DeletedAt string `json:"deleted_at" db:"deleted_at"`
}

type Meta struct {
	TotalData   int64 `json:"total"`
	PerPage     int64 `json:"per_page"`
	TotalPage   int64 `json:"total_page"`
	CurrentPage int64 `json:"current_page"`
}


