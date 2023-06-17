package collectcfg

type Cfg struct {
	MaxCount int32  //触发发送长度
	Path     string //日志位置
}

type Observe struct {
	Host     string
	Username string
	Password string
	Org      string
	Stream   string
}
