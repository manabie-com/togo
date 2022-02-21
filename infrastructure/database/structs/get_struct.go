package structs

type GetStruct struct {
	TableName  string
	Conditions map[string]interface{}
}
