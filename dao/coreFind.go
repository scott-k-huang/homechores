package dao

import (
	"fmt"
	"log"

	"github.com/scott-k-huang/homechores/config"
	"github.com/scott-k-huang/homechores/model"
	"gorm.io/gorm/clause"
)

func FindUser(userId uint) (*model.User, error) {
	db := config.DB
	user := model.User{}
	tx := db.First(&user, userId)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		log.Println("Found User :", user.ID)
		return &user, nil
	}
}

func FindUsers() (*[]model.User, error) {
	db := config.DB
	var users []model.User
	tx := db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		log.Println(fmt.Sprintf("found %v Users", len(users)))
		return &users, nil
	}
}

func FindChoreCategory(choreCategoryId uint) (*model.ChoreCategory, error) {
	db := config.DB
	choreCategory := model.ChoreCategory{}
	tx := db.First(&choreCategory, choreCategoryId)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		log.Println("Found ChoreCategory: ", choreCategory.Name)
		return &choreCategory, nil
	}
}

func FindChoreCategories() (*[]model.ChoreCategory, error) {
	db := config.DB
	var choreCategories []model.ChoreCategory
	tx := db.Find(&choreCategories)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		log.Println(fmt.Sprintf("found %v ChoreCategories", len(choreCategories)))
		return &choreCategories, nil
	}
}

func FindChore(choreId uint) (*model.Chore, error) {
	db := config.DB
	chore := model.Chore{}
	tx := db.Preload(clause.Associations).First(&chore, choreId)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		log.Println("Found Chore: ", chore.Name)
		return &chore, tx.Error
	}
}

func FindChores() (*[]model.Chore, error) {
	db := config.DB
	var chores []model.Chore
	tx := db.Preload("ChoreCategory").Find(&chores)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		log.Println(fmt.Sprintf("found %v Chores", len(chores)))
		return &chores, nil
	}
}
