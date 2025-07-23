package user

import (
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) FindAll() ([]User, error) {
	var users []User
	err := s.DB.Find(&users).Error
	return users, err
}

func (s *UserService) Create(user *User) error {
	return s.DB.Create(user).Error
}
