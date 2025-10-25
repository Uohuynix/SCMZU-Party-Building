package service

import (
	"SCMZU_Party_Building/dao"
	"SCMZU_Party_Building/entity"
)

type TrainingService struct {
	trainingDao *dao.TrainingDAO
}

func NewTrainingService(trainingDao *dao.TrainingDAO) *TrainingService {
	return &TrainingService{trainingDao: trainingDao}
}

func (s *TrainingService) ListTrainings(page, pageSize int) ([]entity.Training, int64, error) {
	return s.trainingDao.FindAll(page, pageSize)
}

func (s *TrainingService) GetTraining(id uint) (*entity.Training, error) {
	return s.trainingDao.FindByID(id)
}

func (s *TrainingService) GetTrainingsByStudent(studentID uint) ([]entity.Training, error) {
	return s.trainingDao.FindByStudentID(studentID)
}

func (s *TrainingService) CreateTraining(training *entity.Training) error {
	return s.trainingDao.Create(training)
}

func (s *TrainingService) UpdateTraining(training *entity.Training) error {
	return s.trainingDao.Update(training)
}

func (s *TrainingService) DeleteTraining(id uint) error {
	return s.trainingDao.Delete(id)
}
