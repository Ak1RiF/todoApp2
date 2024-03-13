package service

import (
	"github.com/todoApp/internal/models"
	"github.com/todoApp/internal/repository"
	"github.com/todoApp/internal/service/dtos"
)

// Quests
type Quests interface {
	GetAll(userId int) []dtos.OutputQuestDto
	GetById(id, userId int) *dtos.OutputQuestDto
	Create(quest models.Quest, userId int)
	Update(id, userId int, quest models.Quest)
	Delete(id, userId int)
}

type Users interface {
	GetUserById(id int) *models.User
	GetUserByUsername(username string) *models.User
	CreateUser(inpUser dtos.InputUserDto)
}

type Service struct {
	Quests Quests
}
type Deps struct {
	Repositories *repository.Repository
}

func NewService(deps Deps) *Service {
	questsService := NewQuestService(deps.Repositories.Quests)
	return &Service{Quests: questsService}
}
