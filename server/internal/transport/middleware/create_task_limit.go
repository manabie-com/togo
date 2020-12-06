package middleware

import (
	"context"
	taskService "github.com/HoangVyDuong/togo/internal/usecase/task"
	userService "github.com/HoangVyDuong/togo/internal/usecase/user"
	"github.com/HoangVyDuong/togo/pkg/define"
	taskDTO "github.com/HoangVyDuong/togo/pkg/dtos/task"
	"github.com/HoangVyDuong/togo/pkg/kit"
	"github.com/HoangVyDuong/togo/pkg/logger"
)

func LimitCreateTask(taskService taskService.Service, userService userService.Service) kit.Middleware{
	return func(endpoint kit.Endpoint) kit.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			logger.Debug("[RateLimit] Start")
			userID, ok := ctx.Value(define.ContextKeyUserID).(string)
			if !ok {
				logger.Error("[RateLimit] Not found userID from context")
				return nil, define.FailedValidation
			}

			isOver := userService.IsOverLimitTask(ctx, userID)
			if isOver {
				logger.Error("[RateLimit] User Over Limit Task")
				return nil, define.UserOverLimitTask
			}

			resp, err := endpoint(ctx, request)
			if err == nil {
				respDTO, ok := resp.(taskDTO.CreateTaskResponse)
				if !ok {
					logger.Debug("[RateLimit] Not Put On CreateTask")
					return nil, define.Unknown
				}

				taskCreated, err := userService.IncreaseTaskTimesPerDuration(ctx, userID, define.RateLimitCreateTaskDuration())
				if err != nil {
					logger.Errorf("[RateLimit] IncreaseTaskTimePerDay In Cache Failed ", err.Error())
					if ok = taskService.Delete(ctx, respDTO.TaskID); !ok {
						logger.Errorf("[RateLimit] Rollback Delete Task In DB Failed ", err.Error())
						return nil, define.Unknown
					}
					return nil, define.Unknown
				}

				if taskCreated > define.RateLimitCreateTaskTimes() {
					logger.Error("[RateLimit] User Over Limit Task")
					if ok = taskService.Delete(ctx, respDTO.TaskID); !ok {
						logger.Errorf("[RateLimit] Rollback Delete Task In DB Failed ", err.Error())
						return nil, define.Unknown
					}
					return nil, define.UserOverLimitTask
				}
			}

			return response, err
		}
	}
}


