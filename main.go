package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TodoTask struct {
	Details   string `json:"details"`
	UserToken string `json:"user_token"`
}

type User struct {
	Token      string `json:"token"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	DailyLimit int    `json:"daily_limit"`
}

type TaskLimitTracker struct {
	UserToken  string `json:"user_token"`
	AddedToday int    `json:"added_today"`
}

var todoTasks []TodoTask
var users []User
var taskLimitTracker []TaskLimitTracker

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	//endpoints and it's handlers
	myRouter.HandleFunc("/", root)
	// myRouter.HandleFunc("/todos", getAllTodos)
	// myRouter.HandleFunc("/users", getAllUsers)
	myRouter.HandleFunc("/todo", createTodoItem).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	//initialize lists
	//added some dummy data for users
	users = []User{
		{Token: "x3EV3uXAVG56w9cn", ID: 1, Name: "James", DailyLimit: 4},
		{Token: "zeXnR3q4VsfgLd3Z", ID: 2, Name: "Richard", DailyLimit: 3},
		{Token: "YJXykGVxTUYh9Gsx", ID: 3, Name: "Jenna", DailyLimit: 5},
	}
	todoTasks = []TodoTask{}
	taskLimitTracker = []TaskLimitTracker{}

	handleRequests()
}

//returns all todo
func getAllTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAllTodos")
	json.NewEncoder(w).Encode(todoTasks)
}

//returs all users
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAllUsers")
	json.NewEncoder(w).Encode(users)
}

//returns a user with the matching token from the api call
func getUserWithToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["token"]

	//call internal function to get user with the same key
	json.NewEncoder(w).Encode(mGetUserWithToken(key))
}

//internal function for getting a user with the same id as the argument
func mGetUserWithId(id int) User {
	for _, user := range users {
		if user.ID == id {
			return user
		}
	}

	return User{}
}

func mGetUserWithToken(token string) User {
	// Loop over all of our users
	// if the user.Token equals the key we pass in
	// return the user encoded as JSON
	for _, user := range users {
		if user.Token == token {
			return user
		}
	}

	return User{}
}

//post handlers
func createTodoItem(w http.ResponseWriter, r *http.Request) {
	//convert the received data to fit the TodoTask object
	reqBody, _ := ioutil.ReadAll(r.Body)
	var todoTask TodoTask
	json.Unmarshal(reqBody, &todoTask)

	//check if user hasn't reached their daily limit. if they haven't, accept their new todo item
	if(mGetUserAddedToday(todoTask.UserToken) < mGetUserWithToken(todoTask.UserToken).DailyLimit){
		todoTasks = append(todoTasks, todoTask)

		//update the user's record on the limit tracker
		mUpdateUserAddLimit(todoTask.UserToken)
	}else{
		//throw an error informing user that their daily limit is reached
		json.NewEncoder(w).Encode("{'message': 'Daily limit reached'}")
	}

	json.NewEncoder(w).Encode(todoTasks)
}

func mUpdateUserAddLimit(userToken string) {
	//when the limit tracker is empty, just append a new record
	if(len(taskLimitTracker) == 0){
		taskLimitTracker = append(taskLimitTracker, TaskLimitTracker{UserToken: userToken, AddedToday: 1})
		fmt.Printf("updateTaskLimit no record, appending\n")
	}else{
		var recordExist bool = false

		//loop through the limit tracker array to check if a record with the same token already exist
		for index, tt := range taskLimitTracker {
			if tt.UserToken == userToken {
				recordExist = true
				
				//if a record is found, increment it's AddedToday field
				taskTracker := &taskLimitTracker[index]
				(*taskTracker).AddedToday = (*taskTracker).AddedToday + 1
				fmt.Printf("updateTaskLimit record found, updating\n")
			}
		}

		//if no record is found, then append a new record for that user
		if(!recordExist){
			taskLimitTracker = append(taskLimitTracker, TaskLimitTracker{UserToken: userToken, AddedToday: 1})
			fmt.Printf("updateTaskLimit record !found, appending\n")
		}
	}

	fmt.Printf("taskLimit: %v\n", taskLimitTracker)
}

//an internal function for getting the AddedToday of a user with the same token
func mGetUserAddedToday(userToken string) int{
	for _, tt := range taskLimitTracker {
		if tt.UserToken == userToken {
			return tt.AddedToday
		}
	}

	return 0
}