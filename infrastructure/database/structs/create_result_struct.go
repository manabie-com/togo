package structs

type CreateResultStruct struct {
	TableName string
	Error     error
	Status    string
	Message   string
	Data      interface{}
}
