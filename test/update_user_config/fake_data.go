package update_user_config

import (
	"math/rand"
	model2 "togo/module/userconfig/model"
)

var UserConfigs = []model2.UserConfig{
	{
		UserId: 1,
		MaxTask: 4,
	},
	{
		UserId: 2,
		MaxTask: 2,
	},
	{
		UserId: 3,
		MaxTask: 7,
	},
	{
		UserId: 4,
		MaxTask: 9,
	},
}

func CreateUserConfigs(length int) []model2.UserConfig {
	data := make([]model2.UserConfig, 0)
	for i := 1; i <= length; i++ {
		id := uint(i)
		maxTask := rand.Intn(100)
		data = append(data, model2.UserConfig{UserId: id,MaxTask: uint(maxTask)})
	}

	return data
}

func CreateUpdatingUserConfigs(length int) []model2.UpdateUserConfig {
	data := make([]model2.UpdateUserConfig, 0)
	for i := 1; i <= length; i++ {
		number := uint(i)
		data = append(data, model2.UpdateUserConfig{MaxTask: &number})
	}

	return data
}