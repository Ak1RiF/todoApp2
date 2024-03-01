package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/todoApp/internal/models"
	"github.com/todoApp/internal/repository"
	"github.com/todoApp/internal/service"
)

type QuestHandler struct {
	questService service.QuestService
}

func NewQuestHandler() *QuestHandler {
	repo := repository.NewRepository()
	deps := service.Deps{Repositories: repo}
	return &QuestHandler{
		questService: *service.NewQuestService(deps.Repositories.Quests),
	}
}

// handlers

func (q *QuestHandler) GetQuests(w http.ResponseWriter, r *http.Request) {
	quests := q.questService.GetAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quests)
}

func (q *QuestHandler) GetQuestById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	questId, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid quest ID", http.StatusBadRequest)
		return
	}

	quest := q.questService.GetById(questId)
	if quest == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quest)
}

func (q *QuestHandler) CreateQuest(w http.ResponseWriter, r *http.Request) {
	var newQuest models.Quest

	err := json.NewDecoder(r.Body).Decode(&newQuest)
	if err != nil {
		http.Error(w, "Invalid body request", http.StatusBadRequest)
		return
	}

	q.questService.Create(newQuest)
	w.WriteHeader(http.StatusCreated)
}

func (q *QuestHandler) UpdateQuest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	questId, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid quest id", http.StatusBadRequest)
		return
	}

	var updateQuest models.Quest

	err = json.NewDecoder(r.Body).Decode(&updateQuest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	q.questService.Update(questId, updateQuest)
	w.WriteHeader(http.StatusOK)
}

func (q *QuestHandler) DeleteQuest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	questId, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid quest ID", http.StatusBadRequest)
		return
	}

	q.questService.Delete(questId)
	w.WriteHeader(http.StatusOK)
}
