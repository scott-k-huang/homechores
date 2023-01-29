package service

import (
	"fmt"
	"github.com/scott-k-huang/homechores/config"
	"github.com/scott-k-huang/homechores/model"
)

func CreateUser(firstName string, lastName string, email string) (*model.User, error) {
	db := config.DB
	var user model.User
	user.FirstName = firstName
	user.LastName = lastName
	user.Email = email
	tx := db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		fmt.Println("Newly created User's ID: ", user.ID)
		return &user, nil
	}
}
