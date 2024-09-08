package api_model

import (
	"time"
)

type Base struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ChoreCategory struct {
	Base
	CreateChoreCategoryRequest
}

type CreateChoreCategoryRequest struct {
	Name string `json:"name"`
}

type Chore struct {
	Base
	Name          string `json:"name"`
	ChoreCategory ChoreCategory
}

type CreateChoreRequest struct {
	Name            string `json:"name"`
	ChoreCategoryID uint   `json:"choreCategoryID"`
}

type User struct {
	Base
	CreateUserRequest
}

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
