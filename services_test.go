package main_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dbSeed(db *sql.DB) error {
	return nil
}

var _ = Describe("Services", func() {
	var db *sql.DB
	var mock sqlmock.Sqlmock
	var service *services.ToDoService

	BeforeEach(func() {
		var err error
		db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		Expect(err).NotTo(HaveOccurred())
		service = &services.ToDoService{
			JWTKey: "wqGyEBBfPK9w3Lxw",
			Store: &sqllite.LiteDB{
				DB: db,
			},
		}
		mock.MatchExpectationsInOrder(false)
		rows := sqlmock.NewRows([]string{"id"}).AddRow("firstUser")
		mock.ExpectQuery(`
	SELECT id FROM users WHERE id = ? AND password = ?
`).WithArgs("firstUser", "example").WillReturnRows(rows)

	})

	Describe("Login entry", func() {

		Context("with right account", func() {
			It("should return http code 200 and token in response body", func() {
				req, err := http.NewRequest("GET", "http://localhost:5050/login?user_id=firstUser&password=example", nil)
				Expect(err).NotTo(HaveOccurred())

				rec := httptest.NewRecorder()
				defer rec.Result().Body.Close()

				service.ServeHTTP(rec, req)
				body, err := ioutil.ReadAll(rec.Body)
				var responseData struct {
					Data string
				}
				json.Unmarshal(body, &responseData)

				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(200))

				Expect(responseData.Data).NotTo(BeEmpty())
			})
		})

		Context("with wrong account", func() {
			It("should return http code 401", func() {
				req, err := http.NewRequest("GET", "http://localhost:5050/login?user_id=wrongUser&password=wrongPassword", nil)
				Expect(err).NotTo(HaveOccurred())

				rec := httptest.NewRecorder()
				defer rec.Result().Body.Close()

				service.ServeHTTP(rec, req)
				body, err := ioutil.ReadAll(rec.Body)
				var responseData struct {
					Error string
				}
				json.Unmarshal(body, &responseData)

				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(401))

				Expect(responseData.Error).To(Equal("incorrect user_id/pwd"))
			})
		})
	})

	Describe("Create todo", func() {
		BeforeEach(func() {
			profileRow := sqlmock.NewRows([]string{"id", "max_todo"}).AddRow("firstUser", 5)
			mock.ExpectQuery("SELECT id, max_todo FROM users WHERE id = ?").
				WithArgs("firstUser").
				WillReturnRows(profileRow)

			mock.ExpectExec(`INSERT INTO tasks (id, content, user_id, created_date) VALUES (?, ?, ?, ?)`).
				WithArgs(sqlmock.AnyArg(), "example task", "firstUser", sqlmock.AnyArg()).
				WillReturnResult(sqlmock.NewResult(1, 1))
		})

		Context("When todos limit not exceeded", func() {
			BeforeEach(func() {
				todoCountRow := sqlmock.NewRows([]string{"count"}).AddRow(1)
				mock.ExpectQuery(`SELECT COUNT(*) FROM tasks WHERE user_id = ? AND created_date >= ?`).
					WithArgs("firstUser", sqlmock.AnyArg()).
					WillReturnRows(todoCountRow)
			})

			It("should success if there is authorized token", func() {
				req, err := http.NewRequest("GET", "http://localhost:5050/login?user_id=firstUser&password=example", nil)
				Expect(err).NotTo(HaveOccurred())

				rec := httptest.NewRecorder()
				defer rec.Result().Body.Close()

				service.ServeHTTP(rec, req)
				body, err := ioutil.ReadAll(rec.Body)
				var responseData struct {
					Data string
				}
				json.Unmarshal(body, &responseData)

				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(200))

				Expect(responseData.Data).NotTo(BeEmpty())

				createTodoBody, err := json.Marshal(map[string]string{
					"content": "example task",
				})
				req, err = http.NewRequest("POST", "http://localhost:5050/tasks", bytes.NewReader(createTodoBody))
				Expect(err).NotTo(HaveOccurred())
				req.Header.Set("Authorization", responseData.Data)
				Expect(err).NotTo(HaveOccurred())

				rec = httptest.NewRecorder()
				defer rec.Result().Body.Close()

				service.ServeHTTP(rec, req)
				body, err = ioutil.ReadAll(rec.Body)
				json.Unmarshal(body, &responseData)

				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(200))

				Expect(responseData.Data).NotTo(BeEmpty())
				var newTodoResponseData struct {
					Data struct {
						ID          string
						Content     string
						UserID      string
						CreatedTime string
					}
				}
				json.Unmarshal(body, &newTodoResponseData)

				Expect(newTodoResponseData.Data.Content).To(Equal("example task"))
			})

			It("should return http code 401 if there is no authorized token", func() {
				createTodoBody, err := json.Marshal(map[string]string{
					"content": "example task",
				})
				req, err := http.NewRequest("POST", "http://localhost:5050/tasks", bytes.NewReader(createTodoBody))
				Expect(err).NotTo(HaveOccurred())
				req.Header.Set("Authorization", "invalid token")
				Expect(err).NotTo(HaveOccurred())

				rec := httptest.NewRecorder()
				defer rec.Result().Body.Close()

				service.ServeHTTP(rec, req)
				body, err := ioutil.ReadAll(rec.Body)
				var responseData struct {
					Data string
				}
				json.Unmarshal(body, &responseData)

				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(401))
			})
		})

		Context("When todos limit exceeed", func() {
			BeforeEach(func() {
				todoCountRow := sqlmock.NewRows([]string{"count"}).AddRow(5)
				mock.ExpectQuery(`SELECT COUNT(*) FROM tasks WHERE user_id = ? AND created_date >= ?`).
					WithArgs("firstUser", sqlmock.AnyArg()).
					WillReturnRows(todoCountRow)
			})

			It("should return status code 451", func() {

				req, err := http.NewRequest("GET", "http://localhost:5050/login?user_id=firstUser&password=example", nil)
				Expect(err).NotTo(HaveOccurred())

				rec := httptest.NewRecorder()
				defer rec.Result().Body.Close()

				service.ServeHTTP(rec, req)
				body, err := ioutil.ReadAll(rec.Body)
				var responseData struct {
					Data string
				}
				json.Unmarshal(body, &responseData)

				Expect(err).NotTo(HaveOccurred())
				Expect(rec.Code).To(Equal(200))

				Expect(responseData.Data).NotTo(BeEmpty())

				createTodoBody, err := json.Marshal(map[string]string{
					"content": "example task",
				})
				req, err = http.NewRequest("POST", "http://localhost:5050/tasks", bytes.NewReader(createTodoBody))
				Expect(err).NotTo(HaveOccurred())
				req.Header.Set("Authorization", responseData.Data)
				Expect(err).NotTo(HaveOccurred())

				rec = httptest.NewRecorder()
				defer rec.Result().Body.Close()

				service.ServeHTTP(rec, req)
				Expect(rec.Code).To(Equal(451))
			})
		})
	})

	Describe("List todos by date", func() {
		BeforeEach(func() {
			todos := sqlmock.NewRows([]string{"ID", "Content", "UserID", "CreatedDate"}).
				AddRow("76c0d5c5-dbfe-47a7-a0a8-dec0ff0d72fb", "task 1", "firstUser", "2021-03-04")
			mock.ExpectQuery(`
			SELECT id, content, user_id, created_date FROM tasks WHERE user_id = ? AND created_date >= ?
			`).
				WithArgs("firstUser", "2021-03-04").
				WillReturnRows(todos)
		})

		It("should return http code 200 with todos list in response body if there is authorized token", func() {
			req, err := http.NewRequest("GET", "http://localhost:5050/login?user_id=firstUser&password=example", nil)
			Expect(err).NotTo(HaveOccurred())

			rec := httptest.NewRecorder()
			defer rec.Result().Body.Close()

			service.ServeHTTP(rec, req)
			body, err := ioutil.ReadAll(rec.Body)
			var responseData struct {
				Data string
			}
			json.Unmarshal(body, &responseData)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(200))

			Expect(responseData.Data).NotTo(BeEmpty())

			req, err = http.NewRequest("GET", "http://localhost:5050/tasks?created_date=2021-03-04", nil)
			Expect(err).NotTo(HaveOccurred())
			req.Header.Set("Authorization", responseData.Data)
			Expect(err).NotTo(HaveOccurred())

			rec = httptest.NewRecorder()
			defer rec.Result().Body.Close()

			service.ServeHTTP(rec, req)
			body, err = ioutil.ReadAll(rec.Body)
			json.Unmarshal(body, &responseData)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(200))

			Expect(responseData.Data).NotTo(BeEmpty())
			var newTodoResponseData struct {
				Data []struct {
					ID          string
					Content     string
					UserID      string `json:"user_id"`
					CreatedDate string `json:"created_date"`
				}
			}
			json.Unmarshal(body, &newTodoResponseData)

			Expect(newTodoResponseData.Data[0].ID).To(Equal("76c0d5c5-dbfe-47a7-a0a8-dec0ff0d72fb"))
			Expect(newTodoResponseData.Data[0].Content).To(Equal("task 1"))
			Expect(newTodoResponseData.Data[0].UserID).To(Equal("firstUser"))
			Expect(newTodoResponseData.Data[0].CreatedDate).To(Equal("2021-03-04"))
		})

		It("should return http code 401 if there is no authorized token", func() {
			req, err := http.NewRequest("GET", "http://localhost:5050/tasks?created_date=2021-03-04", nil)
			Expect(err).NotTo(HaveOccurred())
			req.Header.Set("Authorization", "invalid token")
			Expect(err).NotTo(HaveOccurred())

			rec := httptest.NewRecorder()
			defer rec.Result().Body.Close()

			service.ServeHTTP(rec, req)

			Expect(err).NotTo(HaveOccurred())
			Expect(rec.Code).To(Equal(401))
		})
	})
})
