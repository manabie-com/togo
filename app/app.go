package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/manabie-com/togo/controllers"
	c "github.com/manabie-com/togo/controllers"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	connectURL := os.Getenv("CONNECT_STR")
	db, err := sql.Open("postgres", connectURL)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database is connected!")
	a.DB = db
	a.Router = mux.NewRouter()
	a.Routes()
}

func (a *App) Routes() {

	router := a.Router
	// sub router like http://<HOST>:<PORT>/api/users
	userRouter := router.PathPrefix("/api/users").Subrouter()
	userRouter.HandleFunc("/me", a.GetMe).Methods("GET")
	userRouter.HandleFunc("/signup", a.SignUp).Methods("POST")
	userRouter.HandleFunc("/login", a.Login).Methods("POST")
	userRouter.HandleFunc("/edit", a.UpdateMe).Methods("PATCH")
	// sub router like http://<HOST>:<PORT>/api/tasks
	taskRouter := router.PathPrefix("/api/tasks").Subrouter()
	taskRouter.HandleFunc("", a.GetTasks).Methods("GET")
	taskRouter.HandleFunc("/{id}", a.GetTask).Methods("GET")
	taskRouter.HandleFunc("/add", a.Add).Methods("POST")
	taskRouter.HandleFunc("/{id}", a.Edit).Methods("PATCH")
	taskRouter.HandleFunc("/{id}", a.Delete).Methods("DELETE")
	// sub routes like http://<HOST>:<PORT>/api/payments
	paymentRouter := router.PathPrefix("/api/payments").Subrouter()
	paymentRouter.HandleFunc("", a.Payment).Methods("POST")
	// runs database
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(controllers.JwtAuthentication)
}

func (a *App) GetMe(w http.ResponseWriter, r *http.Request) {
	c.GetTasks(a.DB, w, r)
	c.GetMe(a.DB, w, r)
}

func (a *App) SignUp(w http.ResponseWriter, r *http.Request) {
	c.SignUp(a.DB, w, r)
}

func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	c.Login(a.DB, w, r)
}

func (a *App) UpdateMe(w http.ResponseWriter, r *http.Request) {
	c.UpdateMe(a.DB, w, r)
}

func (a *App) GetTasks(w http.ResponseWriter, r *http.Request) {
	c.GetTasks(a.DB, w, r)
}

func (a *App) GetTask(w http.ResponseWriter, r *http.Request) {
	c.GetTask(a.DB, w, r)
}

func (a *App) Add(w http.ResponseWriter, r *http.Request) {
	c.Add(a.DB, w, r)
}

func (a *App) Edit(w http.ResponseWriter, r *http.Request) {
	c.Edit(a.DB, w, r)
}

func (a *App) Delete(w http.ResponseWriter, r *http.Request) {
	c.Delete(a.DB, w, r)
}

func (a *App) Payment(w http.ResponseWriter, r *http.Request) {
	c.Payment(a.DB, w, r)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
