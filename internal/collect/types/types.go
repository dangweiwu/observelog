package types

import "time"

type Status struct {
	SendFailCount   int32 //发送失败数量
	SendSuccesCount int32 //发送成功数量
	ErrCount        int32 //保存失败条数
	StartTime       time.Time
	LongTime        string
}
