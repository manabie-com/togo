package app

import (
	"net/http"

	c "github.com/manabie-com/togo/controllers"
)

func (a *App) GetMe(w http.ResponseWriter, r *http.Request) {
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

func (a *App) DeleteMe(w http.ResponseWriter, r *http.Request) {
	c.DeleteMe(a.DB, w, r)
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
