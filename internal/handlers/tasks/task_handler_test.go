package tasks

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manabie-com/togo/internal/consts"
	"github.com/manabie-com/togo/internal/models"
	taskService "github.com/manabie-com/togo/internal/services/tasks"
	userService "github.com/manabie-com/togo/internal/services/users"
	"github.com/manabie-com/togo/internal/utils/random"
	requestUtils "github.com/manabie-com/togo/internal/utils/request"
	"github.com/stretchr/testify/assert"
)

type TaskHandlerPostTestCase struct {
	mockUserService   userService.IUserService
	mockTaskService   taskService.ITaskService
	mockCreateRequest ICreateRequest
	requestUserId     string
	expectedRespCode  int
	expectedErr       error
}

func TaskHandlerPostTestCases(t *testing.T) map[string]TaskHandlerPostTestCase {
	return map[string]TaskHandlerPostTestCase{
		"Should create task successfully": {
			mockUserService: userService.MockUserService{
				GetUserByIDResponseUser: &models.User{
					MaxTodo: 5,
				},
				GetUserByIDResponseError: nil,
				ShouldGetUserByIDCalled:  true,
			},
			mockTaskService: taskService.MockTaskService{
				TaskCountByUserResponseCount: int64(random.RandInt(4)),
				TaskCountByUserResponseError: nil,
				ShouldCreateNewCalled:        true,
				ShouldTaskCountByUserCalled:  true,
			},
			mockCreateRequest: MockCreateRequest{
				ShouldBindCalled:     true,
				ShouldValidateCalled: true,
				ShouldToModelCalled:  true,
				BindResponse:         nil,
				ValidateResponse:     nil,
				ToModelResponse: &models.Task{
					UserID:     "test_user_id",
					Content:    "content",
					CreateDate: "date",
				},
			},
			requestUserId:    "test_user_id",
			expectedRespCode: http.StatusOK,
			expectedErr:      nil,
		},
		"Should return invalid request": {
			mockUserService: userService.MockUserService{
				GetUserByIDResponseUser: &models.User{
					MaxTodo: random.RandInt(5),
				},
				GetUserByIDResponseError: nil,
				ShouldGetUserByIDCalled:  true,
			},
			mockTaskService: taskService.MockTaskService{
				TaskCountByUserResponseCount: 6,
				TaskCountByUserResponseError: nil,
				ShouldCreateNewCalled:        true,
				ShouldTaskCountByUserCalled:  true,
			},
			mockCreateRequest: MockCreateRequest{
				ShouldBindCalled:     true,
				ShouldValidateCalled: true,
				ShouldToModelCalled:  true,
				BindResponse:         nil,
				ValidateResponse:     nil,
			},
			requestUserId:    "", // no user id provided
			expectedRespCode: 0,
			expectedErr:      consts.ErrInvalidRequest,
		},
		"Should return max todo reach error": {
			mockUserService: userService.MockUserService{
				GetUserByIDResponseUser: &models.User{
					MaxTodo: random.RandInt(5),
				},
				GetUserByIDResponseError: nil,
				ShouldGetUserByIDCalled:  true,
			},
			mockTaskService: taskService.MockTaskService{
				TaskCountByUserResponseCount: 6,
				TaskCountByUserResponseError: nil,
				ShouldCreateNewCalled:        true,
				ShouldTaskCountByUserCalled:  true,
			},
			mockCreateRequest: MockCreateRequest{
				ShouldBindCalled:     true,
				ShouldValidateCalled: true,
				ShouldToModelCalled:  true,
				BindResponse:         nil,
				ValidateResponse:     nil,
			},
			requestUserId:    random.RandString(10),
			expectedRespCode: 0,
			expectedErr:      consts.ErrMaxTodoReached,
		},
	}
}

func NewMockRequest(method string, body string, userId string) *http.Request {
	bodyByte := []byte(body)
	if body == "" {
		bodyByte = []byte{}
	}
	request, err := http.NewRequest(method, "", bytes.NewReader(bodyByte))
	if err != nil {
		return nil
	}

	if userId != "" {
		request = requestUtils.SetUserID(request, userId)
	}

	return request
}

func TestTaskHandlerPost(t *testing.T) {
	t.Parallel()
	for caseName, tCase := range TaskHandlerPostTestCases(t) {
		t.Run(caseName, func(t *testing.T) {
			response := &httptest.ResponseRecorder{}
			request := NewMockRequest(http.MethodPost, "", tCase.requestUserId)
			taskHandler := TaskHandler{
				NewCreateRequest: func() ICreateRequest { return tCase.mockCreateRequest },
				UserService:      tCase.mockUserService,
				TaskService:      tCase.mockTaskService,
			}
			got := taskHandler.Post(response, request)
			assert.Equal(t, tCase.expectedErr, got)
			assert.Equal(t, response.Code, tCase.expectedRespCode)
		})
	}
}

func TaskHandlerGet(t *testing.T) {
	//TODO: implement
}
