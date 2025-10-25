package service

import (
	"SCMZU_Party_Building/dao"
	"SCMZU_Party_Building/entity"
	"SCMZU_Party_Building/util"
	"errors"
)

type AuthService struct {
	userDao *dao.UserDAO
}

func NewAuthService(userDao *dao.UserDAO) *AuthService {
	return &AuthService{userDao: userDao}
}

func (s *AuthService) Login(username, password string) (*entity.User, error) {
	user, err := s.userDao.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if !util.ComparePassword(user.Password, password) {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

func (s *AuthService) Register(username, password, name, role string) (*entity.User, error) {
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username: username,
		Password: hashedPassword,
		Name:     name,
		Role:     role,
	}

	err = s.userDao.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) GetProfile(userID uint) (*entity.User, error) {
	return s.userDao.FindByID(userID)
}

var ErrInvalidCredentials = errors.New("invalid credentials")
