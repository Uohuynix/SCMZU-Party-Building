package dao

import (
	"SCMZU_Party_Building/entity"

	"gorm.io/gorm"
)

type StudentDAO struct {
	BaseDAO
}

func NewStudentDAO(db *gorm.DB) *StudentDAO {
	return &StudentDAO{BaseDAO: BaseDAO{DB: db}}
}

func (dao *StudentDAO) FindAll(page, pageSize int) ([]entity.Student, int64, error) {
	var students []entity.Student
	var total int64

	offset := (page - 1) * pageSize
	err := dao.DB.Model(&entity.Student{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Offset(offset).Limit(pageSize).Find(&students).Error
	return students, total, err
}

func (dao *StudentDAO) FindByID(id uint) (*entity.Student, error) {
	var student entity.Student
	err := dao.DB.First(&student, id).Error
	return &student, err
}

func (dao *StudentDAO) Create(student *entity.Student) error {
	return dao.DB.Create(student).Error
}

func (dao *StudentDAO) Update(student *entity.Student) error {
	return dao.DB.Save(student).Error
}

func (dao *StudentDAO) Delete(id uint) error {
	return dao.DB.Delete(&entity.Student{}, id).Error
}

func (dao *StudentDAO) FindByStudentNo(studentNo string) (*entity.Student, error) {
	var student entity.Student
	err := dao.DB.Where("student_no = ?", studentNo).First(&student).Error
	return &student, err
}

func (dao *StudentDAO) FindByType(studentType string, page, pageSize int) ([]entity.Student, int64, error) {
	var students []entity.Student
	var total int64

	offset := (page - 1) * pageSize
	err := dao.DB.Model(&entity.Student{}).Where("type = ?", studentType).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Where("type = ?", studentType).Offset(offset).Limit(pageSize).Find(&students).Error
	return students, total, err
}

func (dao *StudentDAO) SearchByName(name string, page, pageSize int) ([]entity.Student, int64, error) {
	var students []entity.Student
	var total int64

	offset := (page - 1) * pageSize
	err := dao.DB.Model(&entity.Student{}).Where("name LIKE ?", "%"+name+"%").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Where("name LIKE ?", "%"+name+"%").Offset(offset).Limit(pageSize).Find(&students).Error
	return students, total, err
}
