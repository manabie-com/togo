package update_user_config

import (
	"context"
	"errors"
	"testing"
	"togo/module/userconfig/model"
	"togo/module/userconfig/repo"
)

type userCfgStore struct {
	Data []model.UserConfig
}

func (u *userCfgStore) Get(ctx context.Context, cond map[string]interface{}) (*model.UserConfig, error) {
	userId := cond["user_id"].(uint)

	for _, usr := range u.Data {
		if usr.UserId == userId {
			return &usr, nil
		}
	}

	return nil, errors.New("DataNotFound")
}

func (u *userCfgStore) Update(ctx context.Context, cond map[string]interface{}, data *model.UpdateUserConfig) error {
	return nil
}

type testCase struct {
	Title string
	UserId uint
	Expect string
}

func TestUpdateUserConfig(t *testing.T)  {
	tcs := []testCase{
		{
			Title: "Update User Config With UserID Is 6",
			UserId: 6,
			Expect: "DataNotFound",
		},
		{
			Title: "Update User Config With UserID Is 1",
			UserId: 1,
			Expect: "",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Title, func(t *testing.T) {
			ctx := context.Background()
			data := CreateUpdatingUserConfigs(1)
			usrCfgStore := &userCfgStore{Data: CreateUserConfigs(6)}
			usrCfgRepo := repo.NewUpdateUserConfigRepo(usrCfgStore)
			if err := usrCfgRepo.UpdateUserConfig(ctx, tc.UserId, &data[0]); err != nil {
				if err.Error() != tc.Expect {
					t.Fatalf("%v", err.Error())
				}
			}
		})
	}
}