package dao

import (
	"SCMZU_Party_Building/entity"
	"gorm.io/gorm"
)

type UserDAO struct {
	BaseDAO
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{BaseDAO: BaseDAO{DB: db}}
}

func (dao *UserDAO) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := dao.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (dao *UserDAO) Create(user *entity.User) error {
	return dao.DB.Create(user).Error
}

func (dao *UserDAO) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	err := dao.DB.First(&user, id).Error
	return &user, err
}

func (dao *UserDAO) Update(user *entity.User) error {
	return dao.DB.Save(user).Error
}
