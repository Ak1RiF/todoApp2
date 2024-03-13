package service

import (
	"log"

	"github.com/todoApp/internal/models"
	"github.com/todoApp/internal/repository"
	"github.com/todoApp/internal/service/dtos"
)

type QuestService struct {
	questRepo repository.Quests
}

func NewQuestService(repo repository.Quests) *QuestService {
	return &QuestService{questRepo: repo}
}

func (s *QuestService) GetAll(userId int) []dtos.OutputQuestDto {
	var quests []dtos.OutputQuestDto

	questsFromDb, err := s.questRepo.Get(userId)

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range questsFromDb {
		questDto := dtos.OutputQuestDto{
			Title:       v.Title,
			Description: v.Description,
			Dificulty:   v.Dificulty,
			Completed:   v.Completed,
		}
		quests = append(quests, questDto)
	}
	return quests
}

func (s *QuestService) GetById(id, userId int) *dtos.OutputQuestDto {
	quest, err := s.questRepo.GetById(id, userId)
	if err != nil {
		log.Fatal(err)
	}
	return &dtos.OutputQuestDto{Title: quest.Title, Description: quest.Description, Dificulty: quest.Dificulty, Completed: quest.Completed}
}

func (s *QuestService) Create(quest models.Quest, userId int) {
	id, err := s.questRepo.Create(quest, userId)
	if id < 0 || err != nil {
		log.Fatal(err)
	}
}

func (s *QuestService) Update(id, userId int, quest models.Quest) {

	err := s.questRepo.Update(id, userId, quest)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *QuestService) Delete(id, userId int) {
	err := s.questRepo.Delete(id, userId)
	if err != nil {
		log.Fatal(err)
	}
}
