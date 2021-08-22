package test

import (
	"context"
	mockRepo "github.com/manabie-com/togo/internal/mocks/storages/repos"
	mockTool "github.com/manabie-com/togo/internal/mocks/tools"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/tools"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
)

func TestLimitTask(t *testing.T) {
	t.Run("Don't have limit task", func(t *testing.T) {
		now := time.Now()
		contextTool := mockTool.IContextTool{}
		contextTool.On("UserIDFromCtx", context.TODO()).Return("1", nil)
		quotaRepo := mockRepo.IQuotaRepo{}
		quotaRepo.On("CountTaskPerDay", context.TODO(), "1", now.Format("2006-01-02")).Return(2, nil)
		quotaRepo.On("GetLimitPerUser", context.TODO(), "1").Return(5, nil)
		quotaService := services.NewQuotaService(&quotaRepo, &contextTool)
		err := quotaService.LimitTask(context.TODO())
		require.Nil(t, err)
	})
	t.Run("Limit task fail by convert context", func(t *testing.T) {
		contextTool := mockTool.IContextTool{}
		errExpect := tools.NewTodoError(500, "fail to convert ctx")
		contextTool.On("UserIDFromCtx", context.TODO()).Return("", errExpect)
		quotaRepo := mockRepo.IQuotaRepo{}
		quotaService := services.NewQuotaService(&quotaRepo, &contextTool)
		err := quotaService.LimitTask(context.TODO())
		require.Error(t, err, errExpect.ErrorMessage)
		require.Equal(t, err, errExpect)
	})
	t.Run("Limit task fail by count task per date", func(t *testing.T) {
		now := time.Now()
		errExpect := tools.NewTodoError(500, "fail to count task in repo")
		contextTool := mockTool.IContextTool{}
		contextTool.On("UserIDFromCtx", context.TODO()).Return("1", nil)
		quotaRepo := mockRepo.IQuotaRepo{}
		quotaRepo.On("CountTaskPerDay", context.TODO(), "1", now.Format("2006-01-02")).Return(0, errExpect)
		quotaService := services.NewQuotaService(&quotaRepo, &contextTool)
		err := quotaService.LimitTask(context.TODO())
		require.Error(t, err, errExpect.ErrorMessage)
		require.Equal(t, err, errExpect)
	})
	t.Run("Limit task fail by get limit by user", func(t *testing.T) {
		now := time.Now()
		errExpect := tools.NewTodoError(500, "fail to get limit in repo")
		contextTool := mockTool.IContextTool{}
		contextTool.On("UserIDFromCtx", context.TODO()).Return("1", nil)
		quotaRepo := mockRepo.IQuotaRepo{}
		quotaRepo.On("CountTaskPerDay", context.TODO(), "1", now.Format("2006-01-02")).Return(2, nil)
		quotaRepo.On("GetLimitPerUser", context.TODO(), "1").Return(0, errExpect)
		quotaService := services.NewQuotaService(&quotaRepo, &contextTool)
		err := quotaService.LimitTask(context.TODO())
		require.Error(t, err, errExpect.ErrorMessage)
		require.Equal(t, err, errExpect)
	})
	t.Run("User reach limit task", func(t *testing.T) {
		now := time.Now()
		errExpect := tools.NewTodoError(http.StatusMethodNotAllowed, "You reach a limit to create task in date")
		contextTool := mockTool.IContextTool{}
		contextTool.On("UserIDFromCtx", context.TODO()).Return("1", nil)
		quotaRepo := mockRepo.IQuotaRepo{}
		quotaRepo.On("CountTaskPerDay", context.TODO(), "1", now.Format("2006-01-02")).Return(5, nil)
		quotaRepo.On("GetLimitPerUser", context.TODO(), "1").Return(5, nil)
		quotaService := services.NewQuotaService(&quotaRepo, &contextTool)
		err := quotaService.LimitTask(context.TODO())
		require.Error(t, err, errExpect.ErrorMessage)
		require.Equal(t, errExpect, err)
	})
}
