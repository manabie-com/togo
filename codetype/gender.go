package codetype

type Gender int

const (
	GenderFemale Gender = 0
	GenderMale   Gender = 1
	GenderOther  Gender = 2
)

func (g *Gender) IsValid() bool {
	switch *g {
	case GenderFemale, GenderMale, GenderOther:
		return true
	default:
		return false
	}
}
