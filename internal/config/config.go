package config

// Config Global Variable
var Config *Conf

type Conf struct {
	AppEnv	string
	AppDebug bool
	AppUrl  string
	AppPort string
	DBHost 	string
	DBPort 	string   
	DBUser 	string
	DBPass 	string
	DBName 	string
}

func InitConfig() {
	// TODO: Load config from .env file
	config := &Conf{
		AppEnv: "local",
		AppDebug: false,
		AppUrl: "localhost",
		AppPort: "5050",
		DBHost: "127.0.0.1",
		DBPort: "5432",
		DBUser: "todo_user",
		DBPass: "todo123",
		DBName: "todo",
	}
	Config = config
}