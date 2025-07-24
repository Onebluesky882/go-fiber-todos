package user

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

// gorm type
type UserService struct {
	DB *gorm.DB
}

// head
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

// func แยก รับ s ของ *UserService
// build-in method FindAll()
// func return []User , error
func (s *UserService) FindAll() ([]User, error) {
	var users []User
	err := s.DB.Find(&users).Error
	return users, err
}

func (s *UserService) Create(input *User) error {
	err := s.DB.Create(input).Error
	if err != nil {
		log.Printf("create user error : %v", err)
	}

	return err
}

func (s *UserService) Find(input *User) (*User, error) {
	var user User
	err := s.DB.Where("email = ?", input.Email).First(&user).Error
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}
	fmt.Printf("%v", err)
	return &user, nil
}
