package model

import "gorm.io/gorm"

type ChoreCategory struct {
	gorm.Model
	Name string
}

type Chore struct {
	gorm.Model
	Name          string
	ChoreCategory ChoreCategory
}

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
}
