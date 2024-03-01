package service

import (
	"log"

	"github.com/todoApp/internal/models"
	"github.com/todoApp/internal/repository"
)

type QuestService struct {
	questRepo repository.Quests
}

func NewQuestService(repo repository.Quests) *QuestService {
	return &QuestService{questRepo: repo}
}

func (s *QuestService) GetAll() []models.Quest {
	quests, err := s.questRepo.Get()
	if err != nil {
		log.Fatal(err)
	}
	return quests
}

func (s *QuestService) GetById(id int) *models.Quest {
	quest, err := s.questRepo.GetById(id)
	if err != nil {
		log.Fatal(err)
	}
	return quest
}

func (s *QuestService) Create(quest models.Quest) {
	id, err := s.questRepo.Create(quest)
	if id < 0 || err != nil {
		log.Fatal(err)
	}
}

func (s *QuestService) Update(id int, quest models.Quest) {

	err := s.questRepo.Update(id, quest)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *QuestService) Delete(id int) {
	err := s.questRepo.Delete(id)
	if err != nil {
		log.Fatal(err)
	}
}
