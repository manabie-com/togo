package snowflake

import (
	"github.com/bwmarrin/snowflake"
)

type snowflakeData struct {}

type SnowflakeData interface {
	GearedID() int
}

func NewSnowflake() SnowflakeData {
	return &snowflakeData{}
}

func (s *snowflakeData) GearedID() int {
	n, err := snowflake.NewNode(1)
	if err != nil {
		return 0
	}
	return int(n.Generate().Int64())
}