package utils

import (
	"github.com/HoangVyDuong/togo/pkg/sonyflake"
)

var sf *sonyflake.Sonyflake

func GenID() uint64 {
	id, _ := sf.NextID()
	return id
}

func init()  {
	sf = sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID:      MachineID,
		CheckMachineID: CheckMachineID,
	})
}

func MachineID() (uint16, error) {
	// for dev environment only
	return 123, nil
}

func CheckMachineID(machineId uint16) bool {
	return true
}
