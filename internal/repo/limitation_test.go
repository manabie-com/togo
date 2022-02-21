package repo_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/chi07/todo/internal/model"
	"github.com/chi07/todo/internal/repo"
)

func TestLimitation_GetByUserID(t *testing.T) {
	if err := util.SetupDB(); err != nil {
		util.Log.Panicf("util.SetupDB(): %v", err)
	}
	defer util.CleanAndClose()
	userID := uuid.MustParse("0a1a88c2-18f9-42dd-b4b0-99ee4dc77751")

	tests := []struct {
		name    string
		userID  uuid.UUID
		wantErr error
		want    *model.Limitation
	}{
		{
			name:   "wrong userID",
			userID: uuid.Nil,
			want:   nil,
		},
		{
			name:   "with valid userID",
			userID: userID,
			want: &model.Limitation{
				ID:        1,
				UserID:    userID.String(),
				LimitTask: 5,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			limitRepo := repo.NewLimitation(util.PostgresDB)
			actual, err := limitRepo.GetByUserID(ctx, tc.userID)
			assert.Equal(t, tc.want, actual)
			assert.NoError(t, err)
		})
	}
}
