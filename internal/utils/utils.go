package utils

import (
	"fmt"
	"lntvan166/togo/internal/repository"
	"net/http"
	"time"
)

const MySQLTimeFormat = "2006-01-02 15:04:05"

func GetCurrentTime() string {
	return time.Now().Format(MySQLTimeFormat)
}

func CheckAccessPermission(w http.ResponseWriter, username string, taskUserID int) error {
	userID, err := repository.GetUserIDByUsername(username)
	if err != nil {
		return err
	}

	if userID != taskUserID {
		return fmt.Errorf("you are not allowed to access this task")
	}

	return nil
}
