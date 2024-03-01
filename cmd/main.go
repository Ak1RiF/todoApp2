package main

import "github.com/todoApp/internal/app"

func main() {
	// repo := repository.NewRepository()
	// deps := service.Deps{Repositories: repo}
	// service := service.NewService(deps)

	// quests := service.Quests.GetAll()
	// fmt.Println(quests)

	// quest := service.Quests.GetById(2)
	// fmt.Println(quest)

	app.Start()
}
