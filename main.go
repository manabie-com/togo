package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/transport"
	"github.com/manabie-com/togo/internal/usecase"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}
func main() {
	pHost := viper.GetString(`database.postgres.host`)
	pPort := viper.GetInt(`database.postgres.port`)
	pUser := viper.GetString(`database.postgres.user`)
	pPass := viper.GetString(`database.postgres.pass`)
	pName := viper.GetString(`database.postgres.name`)
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pHost, pPort, pUser, pPass, pName)
	dbConn, err := sql.Open("postgres", psqlConn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	taskDB := postgres.NewPostgresDB(dbConn)
	taskU := usecase.NewTaskUsecase("", taskDB, timeoutContext)
	tHandler := transport.NewTaskHandler(taskU)
	mux := http.NewServeMux()
	mux.Handle("/", &router{
		handler: tHandler,
	})
	http.ListenAndServe(viper.GetString("server.address"), mux)
}

type router struct {
	handler transport.TaskHandler
}

func (r *router) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")
	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}
	switch req.URL.Path {
	case "/login":
		r.handler.GetAuthToken(resp, req)
		return
	case "/tasks":
		var ok bool
		//req, ok = r.handler.ValidToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch req.Method {
		case http.MethodGet:
			r.handler.ListTasks(resp, req)
		case http.MethodPost:
			r.handler.AddTask(resp, req)
		}
		return
	}
}
