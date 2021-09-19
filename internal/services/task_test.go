package services

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/repositories"
	httpPkg "github.com/manabie-com/togo/pkg/http"
	"github.com/manabie-com/togo/pkg/txmanager"
)

func Test_taskService_ListTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	createdDate := time.Now()
	taskID := "123"
	userID := "firstUser"
	content := "abc"

	ctx := context.Background()
	ctxHasUserId := context.WithValue(ctx, httpPkg.UserIDKey, userID)
	ctxNoUserId := ctx
	ctxHasUserIdEmpty := context.WithValue(ctx, httpPkg.UserIDKey, "")

	type fields struct {
		repo *repositories.Repository
		tx   txmanager.TransactionManager
	}
	type args struct {
		ctx         context.Context
		createdDate time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "reposotory_list_tasks_failed",
			fields: fields{
				repo: &repositories.Repository{
					TaskRepository: func() repositories.TaskRepository {
						repo := repositories.NewMockTaskRepository(ctrl)
						repo.EXPECT().ListTasks(ctxHasUserId, userID, createdDate).Return(
							[]models.Task{
								{
									ID:          taskID,
									Content:     content,
									UserID:      userID,
									CreatedDate: createdDate,
								},
							},
							errors.New("failed"),
						)
						return repo
					}(),
				},
				tx: func() txmanager.TransactionManager {
					return txmanager.NewMockTransactionManager(ctrl)
				}(),
			},
			args: args{
				ctx:         ctxHasUserId,
				createdDate: createdDate,
			},
			wantErr: true,
		},
		{
			name: "list_tasks_successfully",
			fields: fields{
				repo: &repositories.Repository{
					TaskRepository: func() repositories.TaskRepository {
						repo := repositories.NewMockTaskRepository(ctrl)
						repo.EXPECT().ListTasks(ctxHasUserId, userID, createdDate).Return(
							[]models.Task{
								{
									ID:          taskID,
									Content:     content,
									UserID:      userID,
									CreatedDate: createdDate,
								},
							},
							nil,
						)
						return repo
					}(),
				},
				tx: func() txmanager.TransactionManager {
					return txmanager.NewMockTransactionManager(ctrl)
				}(),
			},
			args: args{
				ctx:         ctxHasUserId,
				createdDate: createdDate,
			},
			want: []models.Task{
				{
					ID:          taskID,
					Content:     content,
					UserID:      userID,
					CreatedDate: createdDate,
				},
			},
		},
		{
			name: "context_not_has_user_id",
			fields: fields{
				tx: func() txmanager.TransactionManager {
					return txmanager.NewMockTransactionManager(ctrl)
				}(),
			},
			args: args{
				ctx:         ctxNoUserId,
				createdDate: createdDate,
			},
			wantErr: true,
		},
		{
			name: "context_not_has_user_id_empty",
			fields: fields{
				tx: func() txmanager.TransactionManager {
					return txmanager.NewMockTransactionManager(ctrl)
				}(),
			},
			args: args{
				ctx:         ctxHasUserIdEmpty,
				createdDate: createdDate,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &taskService{
				repo: tt.fields.repo,
				tx:   tt.fields.tx,
			}
			got, err := s.ListTasks(tt.args.ctx, tt.args.createdDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("taskService.ListTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taskService.ListTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskService_AddTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	createdDate := time.Now()
	userId := "firstUser"
	content := "abc"
	taskID := "taskID"
	ctxHasUserId := context.WithValue(ctx, httpPkg.UserIDKey, userId)
	type fields struct {
		repo *repositories.Repository
		tx   txmanager.TransactionManager
	}
	type args struct {
		ctx  context.Context
		task *models.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "add_task_successfully",
			fields: fields{
				tx: func() txmanager.TransactionManager {
					tx := txmanager.NewMockTransactionManager(ctrl)
					tx.EXPECT().Begin(ctxHasUserId).Return(func() txmanager.TransactionManager {
						var childTx = txmanager.NewMockTransactionManager(ctrl)
						childTx.EXPECT().InjectTransaction(ctxHasUserId).Return(ctxHasUserId)
						childTx.EXPECT().End(ctxHasUserId, nil)
						childTx.EXPECT().Recover(ctxHasUserId)
						return childTx
					}())
					return tx
				}(),
				repo: &repositories.Repository{
					TaskRepository: func() repositories.TaskRepository {
						repo := repositories.NewMockTaskRepository(ctrl)
						repo.EXPECT().AddTask(ctxHasUserId, models.Task{
							Content: content,
							UserID:  userId,
						}).Return(taskID, nil)
						repo.EXPECT().GetTask(ctxHasUserId, taskID).Return(
							&models.Task{
								ID:          taskID,
								Content:     content,
								UserID:      userId,
								CreatedDate: createdDate,
							}, nil,
						)
						return repo
					}(),
				},
			},
			args: args{
				ctx: ctxHasUserId,
				task: &models.Task{
					Content: content,
				},
			},
			want: &models.Task{
				ID:          taskID,
				Content:     content,
				UserID:      userId,
				CreatedDate: createdDate,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &taskService{
				repo: tt.fields.repo,
				tx:   tt.fields.tx,
			}
			got, err := s.AddTask(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("taskService.AddTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taskService.AddTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
