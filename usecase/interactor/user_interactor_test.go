package interactor

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/valonekowd/togo/domain/entity"
	"github.com/valonekowd/togo/mocks"
	"github.com/valonekowd/togo/usecase/interfaces"
	"github.com/valonekowd/togo/usecase/request"
	"github.com/valonekowd/togo/usecase/response"
)

func Test_basicUserInteractor_SignIn(t *testing.T) {
	fakePassword := "12345678"
	fakeUser := &entity.User{
		ID:        "this-is-id",
		Email:     "this-is-email",
		Password:  "this-is-password",
		MaxTodo:   3,
		CreatedAt: time.Now().UTC(),
	}
	fakeUser.HashPassword()

	testCases := []struct {
		name          string
		returnUser    *entity.User
		repoError     error
		returnResp    *response.SignIn
		expectedError error
	}{
		{
			name:       "Sign in successful",
			returnUser: fakeUser,
			repoError:  nil,
			returnResp: &response.SignIn{
				Data: &response.SignInPayload{
					AccessToken: "this-is-access-token",
				},
			},
			expectedError: nil,
		},
		{
			name:          "Repository return sql.ErrNoRows",
			returnUser:    nil,
			repoError:     sql.ErrNoRows,
			returnResp:    nil,
			expectedError: fmt.Errorf("finding user: %w", ErrEmailNotFound),
		},
		{
			name:          "Repository return error",
			returnUser:    nil,
			repoError:     errors.New("whoops"),
			returnResp:    nil,
			expectedError: fmt.Errorf("finding user: %w", errors.New("whoops")),
		},
		{
			name:          "Compare password return mismatched error",
			returnUser:    fakeUser,
			repoError:     nil,
			returnResp:    nil,
			expectedError: fmt.Errorf("comparing user password: %w", ErrInvalidEmailOrPassword),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := mocks.UserRepositoryMock{}
			userRepo.On("FindByEmail", mock.MatchedBy(func(ctx context.Context) bool { return true }),
				mock.AnythingOfType("string")).Return(tc.returnUser, tc.repoError)

			userFormatter := mocks.UserFormatterMock{}
			userFormatter.On("SignIn", mock.MatchedBy(func(ctx context.Context) bool { return true }), tc.returnUser).Return(tc.returnResp)

			i := basicUserInteractor{
				ds:        interfaces.DataSource{User: &userRepo},
				presenter: &userFormatter,
			}

			req := request.SignIn{}
			switch tc.name {
			case "Sign in successful":
				req.Password = fakePassword
			case "Compare password return mismatched error":
				req.Password = fakePassword + "wrong-here"
			}

			got, err := i.SignIn(context.Background(), req)
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}

			if got != nil {
				assert.Equal(t, tc.returnResp, got.Data)
			}
		})
	}
}

func Test_basicUserInteractor_SignUp(t *testing.T) {
	fakeUser := &entity.User{
		ID:        "this-is-id",
		Email:     "this-is-email",
		Password:  "this-is-password",
		MaxTodo:   3,
		CreatedAt: time.Now().UTC(),
	}

	testCases := []struct {
		name          string
		user          *entity.User
		repoError     error
		fmtError      error
		returnResp    *response.SignUp
		expectedError error
	}{
		{
			name:      "Sign up successful",
			user:      fakeUser,
			repoError: nil,
			fmtError:  nil,
			returnResp: &response.SignUp{
				Data: &response.SignUpPayload{
					ID:          fakeUser.ID,
					Email:       fakeUser.Email,
					MaxTodo:     fakeUser.MaxTodo,
					CreatedAt:   fakeUser.CreatedAt.Format("2006-02-10 20:00:00 +0700"),
					AccessToken: "this-is-access-token",
				},
			},
			expectedError: nil,
		},
		{
			name:      "Sign up and hash password return error",
			user:      fakeUser,
			repoError: nil,
			fmtError:  nil,
			returnResp: &response.SignUp{
				Data: &response.SignUpPayload{
					ID:          fakeUser.ID,
					Email:       fakeUser.Email,
					MaxTodo:     fakeUser.MaxTodo,
					CreatedAt:   fakeUser.CreatedAt.Format("2006-02-10 20:00:00 +0700"),
					AccessToken: "this-is-access-token",
				},
			},
			expectedError: nil,
		},
		{
			name: "Repository return duplicate email error",
			user: fakeUser,
			repoError: &pq.Error{
				Code:       pq.ErrorCode("23505"),
				Constraint: "users_email_key",
			},
			fmtError:      nil,
			returnResp:    nil,
			expectedError: fmt.Errorf("adding user: %w", ErrEmailAlreadyExists),
		},
		{
			name:          "Repository return error",
			user:          fakeUser,
			repoError:     errors.New("whoops"),
			fmtError:      nil,
			returnResp:    nil,
			expectedError: fmt.Errorf("adding user: %w", errors.New("whoops")),
		},
		{
			name:          "Formatter return error",
			user:          fakeUser,
			repoError:     nil,
			fmtError:      errors.New("whoops"),
			returnResp:    nil,
			expectedError: errors.New("whoops"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := mocks.UserRepositoryMock{}
			userRepo.On("Add", mock.MatchedBy(func(ctx context.Context) bool { return true }),
				mock.MatchedBy(func(t *entity.User) bool { return true })).Return(tc.repoError)

			userFormatter := mocks.UserFormatterMock{}
			userFormatter.On("SignUp", mock.MatchedBy(func(ctx context.Context) bool { return true }),
				mock.MatchedBy(func(t *entity.User) bool { return true })).Return(tc.returnResp, tc.fmtError)

			i := basicUserInteractor{
				ds:        interfaces.DataSource{User: &userRepo},
				presenter: &userFormatter,
			}
			got, err := i.SignUp(context.Background(), request.SignUp{
				Email:    tc.user.Email,
				Password: tc.user.Password,
				MaxTodo:  tc.user.MaxTodo,
			})
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}

			if got != nil {
				assert.Equal(t, tc.returnResp, got)
			}
		})
	}
}
