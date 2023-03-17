package dao

import (
	"github.com/scott-k-huang/homechores/config"
	"gorm.io/gorm"
)

func Paginate(pageNum int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func Save(model gorm.Model) (*gorm.Model, error) {
	db := config.DB
	tx := db.Save(model)
	if tx.Error != nil {
		return nil, tx.Error
	} else {
		return &model, nil
	}
}
