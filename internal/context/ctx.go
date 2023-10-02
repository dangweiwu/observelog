package context

import (
	"gitee.com/lambdang/pkg/logx"
	"github.com/dangweiwu/observelog/internal/config"
)

type AppCtx struct {
	Config config.Config
	Log    *logx.Logx
}

func NewAppCtx(c config.Config) (*AppCtx, error) {
	ctx := &AppCtx{}
	ctx.Config = c
	if lg, err := logx.NewLogx(c.Logx); err != nil {
		return nil, err
	} else {
		ctx.Log = lg
	}

	return ctx, nil
}
