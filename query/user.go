package query

import (
	"manabie-com/togo/entity"
	"strings"
)

func UserByID(id string) (result entity.User, err error) {
	err = entity.Db().Where("id = ?", strings.TrimSpace(id)).First(&result).Error
	return result, err
}

