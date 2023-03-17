package dao

import (
	"github.com/scott-k-huang/homechores/config"
	"github.com/scott-k-huang/homechores/model"
	"log"
)

func CreateUser(firstName string, lastName string, email string) (*model.User, error) {
	db := config.DB
	var user model.User
	log.Println("here is the user %v", user)
	user.FirstName = firstName
	user.LastName = lastName
	user.Email = email
	tx := db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		log.Println("Newly created User's ID: ", user.ID)
		return &user, nil
	}
}

func CreateChoreCategory(name string) (*model.ChoreCategory, error) {
	db := config.DB
	var choreCategory model.ChoreCategory
	choreCategory.Name = name
	tx := db.Create(&choreCategory)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		log.Println("Newly created ChoreCategory ID: ", choreCategory.ID)
		return &choreCategory, nil
	}
}

func UpdateChore(chore model.Chore) (*model.Chore, error) {
	db := config.DB
	tx := db.Save(chore)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		return &chore, nil
	}
}

func CreateChore(name string, choreCategory model.ChoreCategory) (*model.Chore, error) {
	db := config.DB
	var chore model.Chore
	chore.Name = name
	chore.ChoreCategory = choreCategory
	tx := db.Create(&chore)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		log.Println("Newly created Chore ID: ", chore.ID)
		return &chore, nil
	}
}
