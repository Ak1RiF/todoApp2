package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/todoApp/internal/handlers"
)

func Start() {
	router := mux.NewRouter()

	server := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	questHandler := handlers.NewQuestHandler()

	router.HandleFunc("/quests", questHandler.GetQuests).Methods("GET")
	router.HandleFunc("/quests/{id}", questHandler.GetQuestById).Methods("GET")
	router.HandleFunc("/quests", questHandler.CreateQuest).Methods("POST")
	router.HandleFunc("/quests/{id}", questHandler.UpdateQuest).Methods("PUT")
	router.HandleFunc("/quests/{id}", questHandler.DeleteQuest).Methods("DELETE")

	server.ListenAndServe()
}
