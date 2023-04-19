package createtodo_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	createtodo "github.com/manabie-com/togo/internal/features/create_todo"
	"github.com/stretchr/testify/require"
)

func TestCreateTodo(t *testing.T) {
	now := time.Now()

	userID := uuid.New()

	// seed some todos in the previous day
	todos := map[uuid.UUID][]createtodo.Todo{
		userID: {
			{
				ID:          userID,
				Title:       "yesterday title 1",
				Content:     "yesterday content 1",
				UserID:      userID,
				DateCreated: time.Date(now.Year(), now.Month(), now.Day()-1, 1, 0, 0, 0, now.Location()),
				DateUpdated: time.Date(now.Year(), now.Month(), now.Day()-1, 1, 0, 0, 0, now.Location()),
			},
			{
				ID:          uuid.New(),
				Title:       "yesterday title 2",
				Content:     "yesterday content 2",
				UserID:      userID,
				DateCreated: time.Date(now.Year(), now.Month(), now.Day()-1, 2, 0, 0, 0, now.Location()),
				DateUpdated: time.Date(now.Year(), now.Month(), now.Day()-1, 2, 0, 0, 0, now.Location()),
			},
		},
	}

	dailyMaxTodo := 2
	store := newMemStore(todos, dailyMaxTodo)

	feat := createtodo.NewFeature(store)

	ctx := context.Background()

	newTodo1 := createtodo.NewTodo{
		Title:   "title 1",
		Content: "content 1",
		UserID:  userID,
	}

	newTodo2 := createtodo.NewTodo{
		Title:   "title 2",
		Content: "content 2",
		UserID:  userID,
	}

	newTodo3 := createtodo.NewTodo{
		Title:   "title 3",
		Content: "content 3",
		UserID:  userID,
	}

	require := require.New(t)

	todo, err := feat.CreateTodo(ctx, newTodo1)
	require.NoError(err)
	require.Equal(newTodo1, createtodo.NewTodo{
		Title:   todo.Title,
		Content: todo.Content,
		UserID:  todo.UserID,
	})

	todo, err = feat.CreateTodo(ctx, newTodo2)
	require.NoError(err)
	require.Equal(newTodo2, createtodo.NewTodo{
		Title:   todo.Title,
		Content: todo.Content,
		UserID:  todo.UserID,
	})

	todo, err = feat.CreateTodo(ctx, newTodo3) // should error because dailyMaxTodo is 2
	require.ErrorIs(err, createtodo.ErrExceededDailyMaximumTodos)
}

type memStore struct {
	todos        map[uuid.UUID][]createtodo.Todo
	dailyMaxTodo int
	mu           sync.Mutex
}

func newMemStore(todos map[uuid.UUID][]createtodo.Todo, dailyMaxTodo int) *memStore {
	return &memStore{
		todos:        todos,
		dailyMaxTodo: dailyMaxTodo,
	}
}

func (s *memStore) WithinTran(ctx context.Context, fn func(s createtodo.Storer) error) error {
	return fn(s)
}

// CreateTodo inserts a new todo into the database.
func (s *memStore) CreateTodo(ctx context.Context, todo createtodo.Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.todos[todo.UserID] = append(s.todos[todo.UserID], todo)

	return nil
}

// GetUserDailyMaxTodo returns user's daily maximum number of todos allowed.
func (s *memStore) GetUserDailyMaxTodo(ctx context.Context, userID uuid.UUID) (int, error) {
	return s.dailyMaxTodo, nil
}

// GetUserTodayTodoCount returns the total number of today todos of a user.
func (s *memStore) GetUserTodayTodoCount(ctx context.Context, userID uuid.UUID) (int, error) {
	userTodos, ok := s.todos[userID]
	if !ok {
		return 0, nil
	}

	var count int

	for _, todo := range userTodos {
		dateCreated := todo.DateCreated.Local()
		now := time.Now()

		if dateCreated.Year() == now.Year() &&
			dateCreated.Month() == now.Month() &&
			dateCreated.Day() == now.Day() {
			count++
		}
	}

	return count, nil
}
