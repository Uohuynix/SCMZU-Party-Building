package dao

import (
	"gorm.io/gorm"
)

type BaseDAO struct {
	DB *gorm.DB
}

func NewBaseDAO(db *gorm.DB) *BaseDAO {
	return &BaseDAO{DB: db}
}
