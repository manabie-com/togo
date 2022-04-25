package users

import (
	"reflect"
	"testing"
)

func TestCreateTempUsers(t *testing.T) {
	tests := []struct {
		name string
		want []*User
	}{
		{
			name: "Get Temp Users",
			want: []*User{
				{
					UserId:    1,
					Name:      "Test User 1",
					TaskLimit: 5,
					DailyTask: 0,
					TodoTasks: []TodoTask{},
				},
				{
					UserId:    2,
					Name:      "Test User 2",
					TaskLimit: 10,
					DailyTask: 0,
					TodoTasks: []TodoTask{},
				},
				{
					UserId:    3,
					Name:      "Test User 3",
					TaskLimit: 20,
					DailyTask: 0,
					TodoTasks: []TodoTask{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users := CreateTempUsers()
			for index, user := range users {
				if user.UserId != tt.want[index].UserId ||
					user.Name != tt.want[index].Name ||
					user.TaskLimit != tt.want[index].TaskLimit ||
					user.DailyTask != tt.want[index].DailyTask ||
					!reflect.DeepEqual(user.TodoTasks, tt.want[index].TodoTasks) {
					t.Fatalf("Want: %v, got %v", tt.want[index], user)
				}
			}
		})
	}
}
