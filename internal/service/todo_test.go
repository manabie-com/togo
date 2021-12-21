package service_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/api/model"
	"github.com/manabie-com/togo/internal/service"
)

type StubTodoRepo struct {
	Scenario string
}

func (s *StubTodoRepo) Add(m model.Todo) (model.Todo, error) {
	if s.Scenario == "OK" {
		u, _ := uuid.NewUUID()
		m.ID = u.String()
		return m, nil
	} else if s.Scenario == "ERR" {
		return model.Todo{}, service.ErrUnableToAddTodo
	}
	return model.Todo{}, nil
}

func (s *StubTodoRepo) Update(model.Todo) (int, error) {
	return 1, nil
}

func (s *StubTodoRepo) Delete(string) error {
	return nil
}

func (s *StubTodoRepo) GetOne(ID string) (model.Todo, error) {
	return model.Todo{}, nil
}

func (s *StubTodoRepo) GetByUserAndDate(id, date string) ([]model.Todo, error) {
	res, _ := s.Get(id)
	return res, nil
}

func (s *StubTodoRepo) Get(uid string) ([]model.Todo, error) {
	if s.Scenario == "OK" {
		return []model.Todo{
			{ID: "1", Title: "Title", Description: "Desc", CreatedDate: "12-12-2021", UserID: uid},
			{ID: "2", Title: "Title", Description: "Desc", CreatedDate: "12-12-2021", UserID: uid},
		}, nil
	} else if s.Scenario == "EXCEED" {
		return []model.Todo{
			{ID: "1", Title: "Title", Description: "Desc", CreatedDate: "12-12-2021", UserID: uid},
			{ID: "2", Title: "Title", Description: "Desc", CreatedDate: "12-12-2021", UserID: uid},
			{ID: "3", Title: "Title", Description: "Desc", CreatedDate: "12-12-2021", UserID: uid},
			{ID: "4", Title: "Title", Description: "Desc", CreatedDate: "12-12-2021", UserID: uid},
			{ID: "5", Title: "Title", Description: "Desc", CreatedDate: "12-12-2021", UserID: uid},
		}, nil
	}
	return []model.Todo{}, nil
}

func TestAdd(t *testing.T) {

	tcs := []struct {
		scenario string
		err      error
	}{
		{scenario: "OK", err: nil},
		{scenario: "ERR", err: service.ErrUnableToAddTodo},
		{scenario: "EXCEED", err: service.ErrUserExceedDailyTodo},
	}

	for _, v := range tcs {
		tar := &service.DefaultTodo{
			Repo: &StubTodoRepo{
				Scenario: v.scenario,
			},
		}

		res, err := tar.Add(model.Todo{
			Title:       "Test title",
			Description: "Test Desc",
		})

		if v.err == nil && cmp.Equal(res, model.Todo{}) && res.ID != "" && res.CreatedDate != "" {
			t.Fatal("No generated id has been found")
		} else if v.err != err {
			t.Fatalf("Test case failed expected %v but got %v", v.err, err)
		}
	}

}
