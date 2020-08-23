package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"os"

	"github.com/joho/godotenv"

	"github.com/go-pg/pg/v10"

	"github.com/go-pg/pg/v10/orm"

	"github.com/manabie-com/togo/internal/handlers"
	"github.com/manabie-com/togo/internal/middlewares"

	postgre "github.com/manabie-com/togo/internal/storages/postgre"

	repo "github.com/manabie-com/togo/internal/repository"

	entity "github.com/manabie-com/togo/internal/entities"
)

// EnvConfig model
type EnvConfig struct {
	serverHost string
	serverPort string
	dbHost     string
	dbName     string
	dbUser     string
	dbPass     string
	JWTKey     string
}

func main() {
	config := readEnv()
	db := connectDbAndInit(config)

	taskHandler := &handlers.TaskHandler{Repo: &repo.TaskRepository{Store: &postgre.Storage{DB: db}}}
	authHandler := &handlers.AuthHandler{Repo: &repo.UserRepository{Store: &postgre.Storage{DB: db}}, JWTKey: config.JWTKey}

	router := mux.NewRouter()
	router.HandleFunc("/tasks", taskHandler.AddTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.GetByID).Methods("GET", "PUT")
	router.HandleFunc("/tasks", taskHandler.GetAll).Methods("GET").Queries("created_date", "{createdDate}")
	router.Use(middlewares.Authen)
	router.Use(middlewares.LogRequest)
	http.Handle("/", router)

	authRouter := mux.NewRouter()
	authRouter.HandleFunc("/login", authHandler.Login).Methods("POST")
	authRouter.Use(middlewares.LogRequest)
	http.Handle("/login", authRouter)

	log.Printf("Server is listening at %s:%s \n", config.serverHost, config.serverPort)
	http.ListenAndServe(config.serverPort, nil)
}

func connectDbAndInit(config *EnvConfig) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     config.dbHost,
		User:     config.dbUser,
		Password: config.dbPass,
		Database: config.dbName,
	})

	log.Print("Connecting with db \n")

	if db == nil {
		defer db.Close()
		log.Fatal("Error when connect with db. Exit app")
	}

	createSchema(db)

	return db
}

func readEnv() *EnvConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error when reading env. Exit app")
	}

	// temp solution. need to use reflect type to binding value into struct
	config := EnvConfig{
		serverHost: os.Getenv("SERVER_HOST"),
		serverPort: os.Getenv("SERVER_PORT"),
		dbHost:     os.Getenv("DB_HOST"),
		dbUser:     os.Getenv("DB_USER"),
		dbPass:     os.Getenv("DB_PASS"),
		dbName:     os.Getenv("DATABASE"),
		JWTKey:     os.Getenv("JWTKEY"),
	}

	return &config
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*entity.Task)(nil),
		(*entity.User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
