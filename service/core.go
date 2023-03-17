package service

import (
	"github.com/scott-k-huang/homechores/dao"
	"github.com/scott-k-huang/homechores/model"
)

func CreateChore(name string, choreCategoryId uint) (*model.Chore, error) {
	choreCategory, err := dao.FindChoreCategory(choreCategoryId)
	if err != nil {
		return nil, err
	} else {
		chore, err := dao.CreateChore(name, *choreCategory)
		if err != nil {
			return nil, err
		} else {
			return chore, nil
		}
	}
}
