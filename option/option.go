package option

// 控制变量
var Opt Option

type Option struct {
	ConfigPath string   `long:"config" short:"f" description:"配置文件路径" define:"./config.yaml"`
	RunServe   RunServe `command:"run" description:"启动api服务"`
}
