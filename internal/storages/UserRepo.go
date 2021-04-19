//USER//

package storages

import (
	"fmt"
	"log"
)

type IUserRepo interface {
	ValidateUserRepo(user_id, password string) bool	
}

type UserRepo struct {
	IDBHandler
}

func (userRepo *UserRepo) ValidateUserRepo(user_id, password string) bool {

	rows, err := userRepo.Query(fmt.Sprintf("SELECT id FROM users WHERE id = '%v' AND password = '%v'", user_id, password))
	if err != nil {
		return false
	}

	u := User{}
	for rows.Next() {
		err = rows.Scan(&u.ID)
		if err != nil {
			log.Println(err)
			return false
		}
	}

	return true	
}