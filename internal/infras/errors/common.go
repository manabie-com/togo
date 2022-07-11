package errors

var (
	InvalidParamError CustomError = "Invalid param error"
	InvalidIdError    CustomError = "Invalid id, must be integer above 0"
	DataNotFoundError CustomError = "Data does not exist"
)
