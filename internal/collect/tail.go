package collect

import (
	"github.com/dangweiwu/observelog/internal/context"
	"github.com/hpcloud/tail"
	"go.uber.org/zap"
)

type Tailx struct {
	tailCfg tail.Config
	ctx     *context.AppCtx
	msgChan chan string
}

func NewTailx(ctx *context.AppCtx) *Tailx {
	a := &Tailx{
		tailCfg: tail.Config{
			ReOpen:    true,
			Follow:    true,
			Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
			MustExist: false,
			Poll:      true,
		},
		ctx:     ctx,
		msgChan: make(chan string),
	}
	go a.listen()
	return a
}

func (this *Tailx) listen() {
	t, err := tail.TailFile(this.ctx.Config.Collect.Path, this.tailCfg)
	if err != nil {
		this.ctx.Log.Error("启动失败:", zap.String("file", this.ctx.Config.Collect.Path), zap.Error(err))
	} else {
		this.ctx.Log.Info("监听文件", zap.String("file", this.ctx.Config.Collect.Path))
	}
	for {
		select {
		case txt := <-t.Lines:
			if txt.Err != nil {
				this.ctx.Log.Error("采集数据异常", zap.Error(txt.Err))
			}
			if IsJson(txt.Text) {
				this.msgChan <- txt.Text
			} else {
				this.ctx.Log.Error("采集数据非JSON", zap.String("txt", txt.Text))
			}
		}
	}
}

func (this *Tailx) GetMsgChan() chan string {
	return this.msgChan
}
