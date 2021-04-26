package usecase_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/manabie-com/togo/internal/domain"
	"github.com/manabie-com/togo/internal/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/manabie-com/togo/internal/usecase"
	"golang.org/x/crypto/bcrypt"
)

func authTestSetupMockUserStore(t *testing.T, s *mocks.MockUserStore) *mocks.MockUserStore {
	local := make(map[string]domain.User)
	s.EXPECT().CreateUser(gomock.Any()).AnyTimes().DoAndReturn(func(u domain.User) error {
		_, exist := local[u.ID]
		if exist {
			return fmt.Errorf("mock error")
		}
		local[u.ID] = u
		return nil
	})
	s.EXPECT().FindUserByID(gomock.Any()).AnyTimes().DoAndReturn(func(id string) (domain.User, error) {
		u, exist := local[id]
		if !exist {
			return domain.User{}, domain.UserNotFound(id)
		}
		return u, nil
	})
	return s
}

func TestAuthUC(t *testing.T) {
	var authConf = usecase.AuthUCConf{
		Hash: usecase.HashConf{
			Algo:   "bcrypt",
			Cost:   bcrypt.DefaultCost,
			Secret: "", //bcrypt does not use secret
		},
	}
	authConf.Default.MaxTaskPerDay = 5
	c := gomock.NewController(t)
	defer c.Finish()
	mockStore := mocks.NewMockUserStore(c)
	mockStore = authTestSetupMockUserStore(t, mockStore)
	uc, err := usecase.NewAuthUseCase(authConf, mockStore)
	assert.NoError(t, err)
	err = uc.CreateUser("admin", "admin")
	assert.NoError(t, err)

	ok, err := uc.ValidateUser("admin", "admin")
	assert.NoError(t, err)
	assert.True(t, ok)
}
