package service

import (
	"github.com/todoApp/internal/models"
	"github.com/todoApp/internal/repository"
	"github.com/todoApp/internal/service/dtos"
)

// Quests
type Quests interface {
	GetAll() []models.Quest
	GetById(id int) *models.Quest
	Create(quest models.Quest)
	Update(id int, quest models.Quest)
	Delete(id int)
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
