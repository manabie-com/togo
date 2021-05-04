package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/models"
	pkg "github.com/manabie-com/togo/internal/pkg/utils"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/server"
	"testing"
	"time"
)

func init() {
	server.InitServerConfig("../../.env")
	server.Database.InitDatabase()
}

func TestToDoService_GetAuthToken(t *testing.T) {
	type fields struct {
		utils      pkg.Utils
		repository repositories.TaskRepo
	}
	type args struct {
		password string
		userId   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "login_successfully",
			fields: fields{
				utils:      pkg.Utils{},
				repository: repositories.NewTaskRepo(server.Database),
			},
			args: args{
				password: "example",
				userId:   "firstUser",
			},
			want: true,
		},
		{
			name: "login_fail",
			fields: fields{
				utils:      pkg.Utils{},
				repository: repositories.NewTaskRepo(server.Database),
			},
			args: args{
				password: "",
				userId:   "firstUser",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := toDoService{
				utils: tt.fields.utils,
				Repo:  tt.fields.repository,
			}
			exist := a.GetAuthToken(context.Background(), tt.args.userId, tt.args.password)
			if exist != tt.want {
				t.Errorf("Login() got = %v, want %v", exist, tt.want)
			}
		})
	}
}

func TestToDoService_ListTasks(t *testing.T) {
	type fields struct {
		utils      pkg.Utils
		repository repositories.TaskRepo
	}
	type args struct {
		createdDate string
		userId      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_list_tasks_success",
			fields: fields{
				utils:      pkg.Utils{},
				repository: repositories.NewTaskRepo(server.Database),
			},
			args: args{
				createdDate: "2021-05-03",
				userId:      "firstUser",
			},
			want: []*models.Task{
				&models.Task{
					ID:          "11696b55-74f0-4466-a710-2f2221d3042f",
					Content:     "test 1",
					UserID:      "firstUser",
					CreatedDate: "2021-05-03",
				},
				&models.Task{
					ID:          "fdb9d4ce-1ea7-4d65-8dc3-22305420c9de",
					Content:     "test 2",
					UserID:      "firstUser",
					CreatedDate: "2021-05-03",
				},
				&models.Task{
					ID:          "6a06f5f2-76ff-46d6-8841-610038643930",
					Content:     "test 3",
					UserID:      "firstUser",
					CreatedDate: "2021-05-03",
				},
				&models.Task{
					ID:          "68189e08-7689-43f1-82e3-d9b221e5ffe8",
					Content:     "test 4",
					UserID:      "firstUser",
					CreatedDate: "2021-05-03",
				},
				&models.Task{
					ID:          "3f203ee8-f767-405a-8f41-5e6b431e71c7",
					Content:     "test 5",
					UserID:      "firstUser",
					CreatedDate: "2021-05-03",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := toDoService{
				utils: tt.fields.utils,
				Repo:  tt.fields.repository,
			}
			got, err := a.ListTasks(context.Background(), tt.args.userId, tt.args.createdDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("GetListTasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDoService_AddTaskTask(t *testing.T) {
	type fields struct {
		utils      pkg.Utils
		repository repositories.TaskRepo
	}
	type args struct {
		t *models.Task
	}
	uuidStr := uuid.New().String()
	now := time.Now().Format("2006-01-02")
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    error
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_add_task_succeed",
			fields: fields{
				utils:      pkg.Utils{},
				repository: repositories.NewTaskRepo(server.Database),
			},
			args: args{
				t: &models.Task{
					ID:          uuidStr,
					Content:     "test 5",
					UserID:      "firstUser",
					CreatedDate: now,
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := toDoService{
				utils: tt.fields.utils,
				Repo:  tt.fields.repository,
			}
			err := a.AddTask(context.Background(), tt.args.t)
			if err != tt.want {
				t.Errorf("AddTask() got = %v, want %v", err, tt.want)
			}
		})
	}
}
