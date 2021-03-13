package config

// order
// 1. explicit call to Configer.Set
// 2. flag
// 3. env
// 4. config file
// 5. Configer.SetDefault

type Configer interface {
	Load(v interface{}) error
	Set(key string, value interface{})
	SetDefault(key string, value interface{})
}
