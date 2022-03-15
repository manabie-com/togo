package structs

type GetResultStruct struct {
	TableName  string
	Error      error
	Status     string
	Message    string
	Conditions interface{}
	Data       struct {
		RowsAffected int64
		Result       interface{}
	}
}
