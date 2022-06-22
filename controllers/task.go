package controllers

// var GetTasks = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	u.Respond(w, http.StatusOK, map[string]interface{}{})
// }

// var GetTask = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	// check taskID exist
// 	// if not exists
// 	u.Respond(w, http.StatusNotFound, map[string]interface{}{})
// 	// else
// 	u.Respond(w, http.StatusOK, map[string]interface{}{})
// }

// var Add = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	// decode userid from jwt => userId
// 	task := &models.Task{
// 		CreatedAt: time.Now().UTC(),
// 	}
// 	err := json.NewDecoder(r.Body).Decode(task)
// 	if err != nil {
// 		u.Respond(w, http.StatusBadRequest, u.Message(false, "invalid request"))
// 	}

// 	// check task number today greater than or equal to current user limitDayTasks
// 	// if true
// 	u.Respond(w, http.StatusNotAcceptable, map[string]interface{}{})
// 	// else
// 	// update task number today += 1
// 	u.Respond(w, http.StatusCreated, map[string]interface{}{})
// }

// var Edit = func(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// 	// check taskID exist
// 	// if not exists
// 	u.Respond(w, http.StatusNotFound, map[string]interface{}{})
// 	// else
// 	u.Respond(w, http.StatusOK, map[string]interface{}{})
// }
