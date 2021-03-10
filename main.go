package main

import (
	"flag"
	"fmt"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/constant"
	"github.com/manabie-com/togo/internal/rest/api"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	httpPort := flag.Int("port", config.HTTPPort, "which port will be listening")
	flag.Parse()
	srv := api.Rest{
		TodoCtrl: &api.TodoCtrl{
			JWTKey:config.JWT_KEY,
			TodoService: &services.ToDoService{
				UserRepository: storages.GetUserRepository(),
				TaskRepository: storages.GetTaskRepository(),
			},
		},

	}
	hostname, _ := os.Hostname()
	log.Printf("%s :  Starting http at %d", hostname, *httpPort)
	SetupCloseHandler()
	srv.Run(*httpPort)

}

func init() {
	sqllite.InitSqlLiteRepository(constant.SQLITE_DATA_SOURCE)
}

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func SetupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		log.Println("[WARNING] Service Force to Stop")
		os.Exit(0)
	}()
}
