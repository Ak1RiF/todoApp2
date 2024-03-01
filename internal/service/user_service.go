package service

import (
	"log"

	"github.com/todoApp/internal/models"
	"github.com/todoApp/internal/repository"
	"github.com/todoApp/internal/service/dtos"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) GetUserById(id int) *models.User {
	user, err := s.userRepo.GetById(id)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func (s *UserService) GetUserByUsername(username string) *models.User {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func (s *UserService) CreateUser(inpUser dtos.InputUserDto) {
	passwordHash, err := encryptPassword(inpUser.Password)
	if err != nil {
		log.Fatal(err)
	}
	user := models.User{Username: inpUser.Username, PasswordHash: passwordHash}

	id, err := s.userRepo.Create(user)
	if id < 0 || err != nil {
		log.Fatal(err)
	}
}

func encryptPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), err
}
