package repository

import "github.com/todoApp/internal/models"

const (
	databaseUrl = "postgres://postgres:qwerty@localhost:5432/tododb?sslmode=disable"
)

// Quests
type Quests interface {
	Get(userId int) ([]models.Quest, error)
	GetById(id, userId int) (*models.Quest, error)
	Create(quest models.Quest, userId int) (int, error)
	Update(id, userId int, quest models.Quest) error
	Delete(id, userId int) error
}

// Users
type Users interface {
	GetByUsername(username string) (*models.User, error)
	GetById(id int) (*models.User, error)
	Create(user models.User) (int, error)
}

type Repository struct {
	Quests Quests
	Users  Users
}

func NewRepository() *Repository {
	return &Repository{
		Quests: NewQuestRepository(databaseUrl),
		Users:  NewUserRepository(databaseUrl),
	}
}
