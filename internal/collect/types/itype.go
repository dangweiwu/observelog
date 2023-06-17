package types

/*对队列数据进行消费*/
type IDoMsg interface {
	DoMsg([]string) error
}

/*添加mes到任务队列*/
type IDoOneMsg interface {
	DoOneMsg(msg string) error
}

// 采集数据接口
type ICollect interface {
	Collect() chan string
}

// 发送数据
type ISend interface {
	Send(string) (oknum, failed int, err error)
}
