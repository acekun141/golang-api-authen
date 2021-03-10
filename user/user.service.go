package user

import (
	"errors"
	"learn-gin/util"

	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func (s *Service) GetAllUser() map[string]interface{} {
	result := map[string]interface{}{}
	s.DB.Model(&User{}).Take(result)
	return result
}

func (s *Service) GetUserLogin(user User) (User, error) {
	var existedUser User
	if tx := s.DB.Model(&User{}).First(&existedUser, User{Username: user.Username}); tx.Error != nil {
		return User{}, errors.New("Username or Password incorrect")
	}
	valid := util.ComparePassword(user.Password, existedUser.Password)
	if valid {
		return existedUser, nil
	}
	return User{}, errors.New("Username or Password incorrect")
}

func (s *Service) CreateUser(user *User) error {
	var existedUser User
	if tx := s.DB.Model(&User{}).First(&existedUser, User{Username: user.Username}); tx.Error != nil {
		hashPassword, err := util.HashAndSalt(user.Password)
		if err != nil {
			return errors.New("Server Error")
		}
		s.DB.Model(&User{}).Create(&User{
			Username: user.Username,
			Email:    user.Email,
			Password: hashPassword,
		})
		return nil
	}
	return errors.New("User existed")
}
