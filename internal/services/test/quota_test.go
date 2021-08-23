package test

import (
	"context"
	"net/http"
	"testing"
	"time"

	mockRepo "github.com/manabie-com/togo/internal/mocks/storages"
	mockTool "github.com/manabie-com/togo/internal/mocks/tools"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/tools"
	"github.com/stretchr/testify/require"
)

func TestLimitTask(t *testing.T) {
	t.Run("Don't have limit task", func(t *testing.T) {
		now := time.Now()
		contextTool := mockTool.IContextTool{}
		contextTool.On("UserIDFromCtx", context.TODO()).Return("1", nil)
		quotaRepo := mockRepo.IQuotaRepo{}
		quotaRepo.On("CountTaskPerDayStore", context.TODO(),
			storages.CountTaskPerDayParams{UserID: "1", CreatedDate: now.Format("2006-01-02")}).
			Return(int64(2), nil)
		quotaRepo.On("GetLimitPerUserStore", context.TODO(), "1").Return(int32(5), nil)
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
		quotaRepo.On("CountTaskPerDayStore", context.TODO(),
			storages.CountTaskPerDayParams{UserID: "1", CreatedDate: now.Format("2006-01-02")}).
			Return(int64(0), errExpect)
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
		quotaRepo.On("CountTaskPerDayStore", context.TODO(),
			storages.CountTaskPerDayParams{UserID: "1", CreatedDate: now.Format("2006-01-02")}).
			Return(int64(2), nil)
		quotaRepo.On("GetLimitPerUserStore", context.TODO(), "1").Return(int32(0), errExpect)
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
		quotaRepo.On("CountTaskPerDayStore", context.TODO(),
			storages.CountTaskPerDayParams{UserID: "1", CreatedDate: now.Format("2006-01-02")}).
			Return(int64(5), nil)
		quotaRepo.On("GetLimitPerUserStore", context.TODO(), "1").Return(int32(5), nil)
		quotaService := services.NewQuotaService(&quotaRepo, &contextTool)
		err := quotaService.LimitTask(context.TODO())
		require.Error(t, err, errExpect.ErrorMessage)
		require.Equal(t, errExpect, err)
	})
}
