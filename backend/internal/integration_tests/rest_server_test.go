package integration_tests

import (
	"testing"
	"go.uber.org/dig"
	"manabie.com/internal/common"
	"manabie.com/internal/repositories"
	"manabie.com/internal/controllers"
	"manabie.com/internal/views"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"database/sql"
	"encoding/json"
	"bytes"
)

func TestViewRest(t *testing.T) {
	container := dig.New()

	var err error
	err = common.ProvideClockSim(container)
	if err != nil {
		t.Fatal(err)
	}
	err = repositories.ProvideSqlConnection(container)
	if err != nil {
		t.Fatal(err)
	}
	err = repositories.ProvideUserRepository(container)
	if err != nil {
		t.Fatal(err)
	}
	err = repositories.ProvideTaskRepository(container)
	if err != nil {
		t.Fatal(err)
	}
	err = repositories.ProvideRepositoryFactory(container)
	if err != nil {
		t.Fatal(err)
	}
	err = controllers.ProvideTaskController(container)
	if err != nil {
		t.Fatal(err)
	}
	err = views.ProvideTaskViewRest(container)
	if err != nil {
		t.Fatal(err)
	}

	container.Invoke(func (iView *views.TaskViewRest) {
		http.Handle("/tasks", iView)
	})


	t.Run("task view", func (t *testing.T) {
		//// add user to database
		err := container.Invoke(func (iDb *sql.DB) {
			iDb.Query(`DELETE FROM "tasks"`)
			iDb.Query(`DELETE FROM "user"`)
			iDb.Query(`INSERT INTO "user" (
				id,
				name,
				task_limit
			) VALUES (
				1,
				'user-1',
				2
			)`)
		})
		if err != nil {
			t.Fatal(err)
		}
		err = container.Invoke(func (iView *views.TaskViewRest, ) {
			server := httptest.NewServer(iView)
			client := server.Client()

			cookie := &http.Cookie{
				Name:   "user_id",
				Value:  "1",
				MaxAge: 300,
			}
			requestBody := map[string]interface{} {
				"title": "abc",
				"content": "def",
			}
			requestBodyJson, err := json.Marshal(requestBody)
			req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(requestBodyJson))
			if err != nil {
				t.Fatal(err)
			}
			req.AddCookie(cookie)

			_, err = client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
		})
		if err != nil {
			t.Fatal(err)
		}

		err = container.Invoke(func (iDb *sql.DB) {
			result, err := iDb.Query(`SELECT 
				id,
				title,
				content,
				created_time
			FROM "tasks" WHERE user_id=1`)
			if err != nil {
				t.Fatal(err)
			}
			defer result.Close()
			count := 0
			for result.Next() {
				var id int
				var title, content string
				var createdTime common.Time
				result.Scan(
					&id,
					&title,
					&content,
					&createdTime,
				)
				require.NotEqual(t, -1, id)
				require.Equal(t, "abc", title)
				require.Equal(t, "def", content)

				count += 1
			}
			require.Equal(t, 1, count)
		})
		if err != nil {
			t.Fatal(err)
		}
	})
}
