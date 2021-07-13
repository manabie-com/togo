package utils

const (
	LimitDefault = 10
)

func GetLimitOffsetFormPageNumber(page, limit int) (int, int) {
	if limit <= 0 {
		limit = LimitDefault
	}
	if page < 0 {
		page = 0
	}
	offset := limit * page
	return limit, offset
}
