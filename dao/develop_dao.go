package dao

import (
	"SCMZU_Party_Building/entity"
	"gorm.io/gorm"
)

type DevelopmentDAO struct {
	BaseDAO
}

func NewDevelopmentDAO(db *gorm.DB) *DevelopmentDAO {
	return &DevelopmentDAO{BaseDAO: BaseDAO{DB: db}}
}

func (dao *DevelopmentDAO) FindByStudentID(studentID uint) (*entity.Development, error) {
	var development entity.Development
	err := dao.DB.Where("student_id = ?", studentID).Preload("Student").First(&development).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &development, err
}

func (dao *DevelopmentDAO) Create(development *entity.Development) error {
	return dao.DB.Create(development).Error
}

func (dao *DevelopmentDAO) Update(development *entity.Development) error {
	return dao.DB.Save(development).Error
}
