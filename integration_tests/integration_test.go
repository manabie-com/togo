package integration_tests

import (
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	db, err := initDB()
	if err != nil {
		t.Error(err)
	}

	// test login
	ts := prepareServer(db)
	defer ts.Close()

	type args struct {
		userID   string
		password string
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
	}{
		{
			name: "valid user",
			args: args{
				userID:   "firstUser",
				password: "example",
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "invalid user",
			args: args{
				userID:   "invalidUser",
				password: "example",
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "incorrect password",
			args: args{
				userID:   "firstUser",
				password: "incorrectPassword",
			},
			wantStatus: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := generateLoginRequest(ts.URL+"/login", tt.args.userID, tt.args.password)
			if err != nil {
				t.Error(err)
				return
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Error(err)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode != tt.wantStatus {
				t.Errorf("status got %d, want %d", resp.StatusCode, tt.wantStatus)
				return
			}
		})
	}
}

func TestAddTask(t *testing.T) {
	db, err := initDB()
	if err != nil {
		t.Error(err)
	}

	ts := prepareServer(db)
	defer ts.Close()

	token, err := login(ts.URL+"/login", "firstUser", "example")
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("create task successfully", func(t *testing.T) {
		addTaskRequest, err := generateAddTaskRequest(ts.URL+"/tasks", token)
		if err != nil {
			t.Error(err)
			return
		}
		resp, err := http.DefaultClient.Do(addTaskRequest)
		if err != nil {
			t.Error(err)
			return
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("invalid status code %d, want %d", resp.StatusCode, http.StatusOK)
			return
		}
	})
	t.Run("create task unsuccessfully", func(t *testing.T) {
		addTaskRequest, err := generateAddTaskRequest(ts.URL+"/tasks", "invalid token")
		if err != nil {
			t.Error(err)
			return
		}
		resp, err := http.DefaultClient.Do(addTaskRequest)
		if err != nil {
			t.Error(err)
			return
		}
		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("invalid status code %d, want %d", resp.StatusCode, http.StatusOK)
			return
		}
	})
}

func TestRetrieveTasks(t *testing.T) {
	db, err := initDB()
	if err != nil {
		t.Error(err)
	}

	ts := prepareServer(db)
	defer ts.Close()

	token, err := login(ts.URL+"/login", "firstUser", "example")
	if err != nil {
		t.Error(err)
		return
	}

	numberOfTasks := 5
	addedTasks := 0
	for i := 0; i < numberOfTasks; i++ {
		addTaskRequest, err := generateAddTaskRequest(ts.URL+"/tasks", token)
		if err != nil {
			t.Error(err)
			return
		}
		_, err = http.DefaultClient.Do(addTaskRequest)
		if err == nil {
			addedTasks++
		}
	}

	t.Run("retrieve tasks successfully", func(t *testing.T) {
		tasks, err := retrieveTasks(ts.URL+"/tasks", token)
		if err != nil {
			t.Error(err)
			return
		}
		if len(tasks) != addedTasks {
			t.Errorf("number of added tasks (%d) doesn't equal number of retrieved tasks (%d)", addedTasks, len(tasks))
			return
		}
	})

	t.Run("retrieve tasks unsuccessfully", func(t *testing.T) {
		_, err := retrieveTasks(ts.URL+"/tasks", "fake_token")
		if err == nil {
			t.Error("error expected")
			return
		}
	})
}

func TestLimitReached(t *testing.T) {
	db, err := initDB()
	if err != nil {
		t.Error(err)
	}

	ts := prepareServer(db)
	defer ts.Close()

	token, err := login(ts.URL+"/login", "firstUser", "example")
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("limit reached", func(t *testing.T) {
		maxTasks := 5
		numberOfTasks := 6
		addedTasks := 0
		for i := 0; i < numberOfTasks; i++ {
			addTaskRequest, err := generateAddTaskRequest(ts.URL+"/tasks", token)
			if err != nil {
				t.Error(err)
				return
			}
			addedTasks++
			resp, err := http.DefaultClient.Do(addTaskRequest)
			if err != nil {
				t.Error(err)
				return
			}
			if addedTasks <= maxTasks {
				if resp.StatusCode != http.StatusOK {
					t.Errorf("Status is not valid %d, added tasks %d, maximum %d", resp.StatusCode, addedTasks, maxTasks)
					return
				}
			} else {
				if resp.StatusCode == http.StatusOK {
					t.Errorf("Status is not valid %d, added tasks %d, maximum %d", resp.StatusCode, addedTasks, maxTasks)
					return
				}
			}
		}
	})
}
