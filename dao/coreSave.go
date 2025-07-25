package dao

import (
	"log"

	"github.com/scott-k-huang/homechores/config"
	"github.com/scott-k-huang/homechores/model"
)

func CreateUser(firstName string, lastName string, email string) (*model.User, error) {
	db := config.DB
	var user model.User
	log.Printf("here is the user %v", user)
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

func UpdateUser(user model.User) (*model.User, error) {
	db := config.DB
	tx := db.Save(user)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
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

func UpdateChoreCategory(choreCategory model.ChoreCategory) (*model.ChoreCategory, error) {
	db := config.DB
	tx := db.Save(choreCategory)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		return &choreCategory, nil
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

func UpdateChore(chore model.Chore) (*model.Chore, error) {
	db := config.DB
	tx := db.Save(chore)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		return &chore, nil
	}
}
