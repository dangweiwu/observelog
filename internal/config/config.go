package config

import (
	"gitee.com/lambdang/pkg/logx"
	"observelog/internal/collect/collectcfg"
)

type Config struct {
	Host     string `default:"127.0.0.1:8000"`
	User     string `default:"root"` //web账号
	Password string `default:"root"` //web密码
	Logx     logx.LogxConfig
	Observe  collectcfg.Observe
	Collect  collectcfg.Cfg
}
