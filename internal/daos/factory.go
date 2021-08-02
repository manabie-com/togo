package daos

import "log"

var (
	accountDAO = AccountDAO{}
	taskDAO    = TaskDAO{}
)

func init() {
	log.Println("Initializing DAO Factory")

}

func GetAccountDAO() AccountDAO {
	return accountDAO
}

func GetTaskDAO() TaskDAO {
	return taskDAO
}
