package main

import (
	"github.com/dangweiwu/observelog/option"
	flags "github.com/jessevdk/go-flags"
)

/*
1. 能采集信息发送到指定位置
2. 能监控当前状态
*/

func main() {
	p := flags.NewParser(&option.Opt, flags.Default)
	p.ShortDescription = "v1.0 log observe"
	p.LongDescription = `v1.0 log observe`
	p.Parse()
}
