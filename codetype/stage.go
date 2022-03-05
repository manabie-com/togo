package codetype

import "strings"

type StageType string

const (
	StageTypeProd  StageType = "PROD"
	StageTypeDev   StageType = "DEV"
	StageTypeLocal StageType = "LOCAL"
)

func (s *StageType) UpCase() {
	*s = StageType(strings.ToUpper(string(*s)))
}

func (s *StageType) ToString() string {
	return string(*s)
}

func (s *StageType) IsValid() bool {
	switch *s {
	case StageTypeProd, StageTypeDev, StageTypeLocal:
		return true
	default:
		return false
	}
}

func (s *StageType) IsProd() bool {
	return *s == StageTypeProd
}

func (s *StageType) IsDev() bool {
	return *s == StageTypeDev
}

func (s *StageType) IsLocal() bool {
	return *s == StageTypeLocal
}
