package dao

import (
	"SCMZU_Party_Building/entity"

	"gorm.io/gorm"
)

type RewardDAO struct {
	BaseDAO
}

func NewRewardDAO(db *gorm.DB) *RewardDAO {
	return &RewardDAO{BaseDAO: BaseDAO{DB: db}}
}

func (dao *RewardDAO) FindByStudentID(studentID uint) ([]entity.Reward, error) {
	var rewards []entity.Reward
	err := dao.DB.Where("student_id = ?", studentID).Preload("Student").Find(&rewards).Error
	return rewards, err
}

func (dao *RewardDAO) Create(reward *entity.Reward) error {
	return dao.DB.Create(reward).Error
}

func (dao *RewardDAO) Update(reward *entity.Reward) error {
	return dao.DB.Save(reward).Error
}

func (dao *RewardDAO) Delete(id uint) error {
	return dao.DB.Delete(&entity.Reward{}, id).Error
}

func (dao *RewardDAO) FindByType(rewardType string, page, pageSize int) ([]entity.Reward, int64, error) {
	var rewards []entity.Reward
	var total int64

	offset := (page - 1) * pageSize
	err := dao.DB.Model(&entity.Reward{}).Where("type = ?", rewardType).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Where("type = ?", rewardType).Offset(offset).Limit(pageSize).Preload("Student").Find(&rewards).Error
	return rewards, total, err
}

func (dao *RewardDAO) FindAll(page, pageSize int) ([]entity.Reward, int64, error) {
	var rewards []entity.Reward
	var total int64

	offset := (page - 1) * pageSize
	err := dao.DB.Model(&entity.Reward{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Offset(offset).Limit(pageSize).Preload("Student").Find(&rewards).Error
	return rewards, total, err
}
