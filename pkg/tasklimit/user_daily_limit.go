package tasklimit

import "math/rand"

type UserLimitSvc interface {
	GetUserLimit(userID uint64) uint32
}

type userLimitCache struct {
	userLimit map[uint64]uint32
}

func (u *userLimitCache) GetUserLimit(userID uint64) uint32 {
	limit, ok := u.userLimit[userID]
	if !ok {
		newLimit := uint32(rand.Intn(5))
		u.userLimit[userID] = newLimit
		return newLimit
	}

	return limit
}

func GetUserLimiSvc() UserLimitSvc {
	return &userLimitCache{userLimit: make(map[uint64]uint32)}
}
