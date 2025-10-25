package dao

import (
	"SCMZU_Party_Building/entity"

	"gorm.io/gorm"
)

type TrainingDAO struct {
	BaseDAO
}

func NewTrainingDAO(db *gorm.DB) *TrainingDAO {
	return &TrainingDAO{BaseDAO: BaseDAO{DB: db}}
}

func (dao *TrainingDAO) FindAll(page, pageSize int) ([]entity.Training, int64, error) {
	var trainings []entity.Training
	var total int64

	offset := (page - 1) * pageSize
	err := dao.DB.Model(&entity.Training{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Offset(offset).Limit(pageSize).Preload("Student").Find(&trainings).Error
	return trainings, total, err
}

func (dao *TrainingDAO) FindByID(id uint) (*entity.Training, error) {
	var training entity.Training
	err := dao.DB.Preload("Student").First(&training, id).Error
	return &training, err
}

func (dao *TrainingDAO) FindByStudentID(studentID uint) ([]entity.Training, error) {
	var trainings []entity.Training
	err := dao.DB.Where("student_id = ?", studentID).Preload("Student").Find(&trainings).Error
	return trainings, err
}

func (dao *TrainingDAO) Create(training *entity.Training) error {
	return dao.DB.Create(training).Error
}

func (dao *TrainingDAO) Update(training *entity.Training) error {
	return dao.DB.Save(training).Error
}

func (dao *TrainingDAO) Delete(id uint) error {
	return dao.DB.Delete(&entity.Training{}, id).Error
}

func (dao *TrainingDAO) FindByPeriod(period string, page, pageSize int) ([]entity.Training, int64, error) {
	var trainings []entity.Training
	var total int64

	offset := (page - 1) * pageSize
	err := dao.DB.Model(&entity.Training{}).Where("period = ?", period).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Where("period = ?", period).Offset(offset).Limit(pageSize).Preload("Student").Find(&trainings).Error
	return trainings, total, err
}

func (dao *TrainingDAO) FindByScore(score string, page, pageSize int) ([]entity.Training, int64, error) {
	var trainings []entity.Training
	var total int64

	offset := (page - 1) * pageSize
	err := dao.DB.Model(&entity.Training{}).Where("score = ?", score).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Where("score = ?", score).Offset(offset).Limit(pageSize).Preload("Student").Find(&trainings).Error
	return trainings, total, err
}
