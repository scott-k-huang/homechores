package api_model

import "time"

type Base struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ChoreCategory struct {
	Base
	Name string `json:"name"`
}

type Chore struct {
	Base
	Name          string `json:"name"`
	ChoreCategory ChoreCategory
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
