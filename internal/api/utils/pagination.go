package utils

const (
	LimitDefault = 10
)

func GetLimitOffsetFormPageNumber(page, limit int) (int, int) {
	if limit <= 0 {
		limit = LimitDefault
	}
	offset := limit * page
	return limit, offset
}
