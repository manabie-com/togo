package structs

type GetResultStruct struct {
	TableName  string
	Error      error
	Status     string
	Message    string
	Conditions map[string]interface{}
	Data       interface{}
}
