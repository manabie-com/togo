package auth_services

import (
	"mini_project/db/model"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	unassign = 0
	monitor  = 1
	operator = 2
	admin    = 3
)

var (
	roleLevel = map[string]int{
		"unassign": unassign,
		"monitor":  monitor,
		"operator": operator,
		"admin":    admin,
	}
)

func ConvertTime2Timestamppb(t *time.Time) *timestamppb.Timestamp {
	if t != nil {
		timepb, err := ptypes.TimestampProto(*t)
		if err != nil {
			panic(err)
		}
		return timepb
	}
	return nil
}

func appendMap(a, b map[string]string) map[string]string {
	for k, v := range b {
		a[k] = v
	}
	return a
}

func ValidateCreateUserRequest(req *UserRequest, db model.DatabaseAPI) map[string]string {
	// validate email
	var detailError = map[string]string{}

	if !db.IsUserNotExists(req.Name) {
		detailError["name"] = "This ID cannot be used"
	}

	return nil
}
