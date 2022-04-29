package main

import (
	"fmt"
	"net/http"
	"manabie.com/internal/common"
	"manabie.com/internal/repositories"
	"manabie.com/internal/controllers"
	"manabie.com/internal/views"
	"go.uber.org/dig"
	"log"
)

func main() {
	container := dig.New()

	var err error
	err = common.ProvideClockSim(container)
	if err != nil {
		log.Fatal(err)
	}
	err = repositories.ProvideSqlConnection(container)
	if err != nil {
		log.Fatal(err)
	}
	err = repositories.ProvideUserRepository(container)
	if err != nil {
		log.Fatal(err)
	}
	err = repositories.ProvideTaskRepository(container)
	if err != nil {
		log.Fatal(err)
	}
	err = repositories.ProvideRepositoryFactory(container)
	if err != nil {
		log.Fatal(err)
	}
	err = controllers.ProvideTaskController(container)
	if err != nil {
		log.Fatal(err)
	}
	err = views.ProvideTaskViewRest(container)
	if err != nil {
		log.Fatal(err)
	}

	container.Invoke(func (iView *views.TaskViewRest) {
		http.Handle("/tasks", iView)
	})

	fmt.Println("Http server started")
	// http.HandleFunc("/tasks", views.TaskViewRest)
	http.ListenAndServe(":8090", nil)
}