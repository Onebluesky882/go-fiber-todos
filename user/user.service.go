package user

import (
	"log"

	"github.com/google/uuid"
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

func (s *UserService) FindAll() ([]User, error) {
	var users []User
	err := s.DB.Find(&users).Error
	return users, err
}

func (s *UserService) Create(input *User) error {
	if input.ID == uuid.Nil {
		input.ID = uuid.New()
	}
	err := s.DB.Create(input).Error
	if err != nil {
		log.Printf("create user error : %v", err)
	}

	return err
}
