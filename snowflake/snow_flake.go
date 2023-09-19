package snowflake

import (
	"fmt"

	"github.com/sony/sonyflake"
)

// 雪花ID工具

var _sf *sonyflake.Sonyflake

func InitSnowflake(machineId uint16) {
	if _sf != nil {
		return
	}

	var st sonyflake.Settings
	st.MachineID = func() (u uint16, e error) {
		return machineId, nil
	}
	_sf = sonyflake.NewSonyflake(st)
}

func NextId() (uint64, error) {
	id, err := _sf.NextID()
	if err != nil {
		fmt.Errorf("snowflake NextId ：%s", err.Error())
	}
	return id, nil
}
