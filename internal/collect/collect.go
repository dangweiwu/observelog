package collect

import (
	"fmt"
	"github.com/dangweiwu/observelog/internal/collect/types"
	"github.com/dangweiwu/observelog/internal/context"
	"github.com/hpcloud/tail"
	"go.uber.org/zap"
	"strings"
	"sync/atomic"
	"time"
)

/*
collect tail

*/

type Collect struct {
	tailcfg         tail.Config
	ctx             *context.AppCtx
	MsgChan         chan string
	MsgTxt          strings.Builder
	MsgCount        int32
	act             chan struct{}
	sender          types.ISend
	startTime       time.Time
	SendSuccesCount int32 //发送成功次数
	SendFailCount   int32 //发送失败次数
	ErrCount        int32 //处理失败条数
	ErrMsg          atomic.Value
}

func NewCollect(ctx *context.AppCtx) *Collect {
	a := &Collect{tailcfg: tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}, ctx: ctx, MsgTxt: strings.Builder{}, startTime: time.Now(), act: make(chan struct{})}
	return a

}

func (this *Collect) SetMsgChan(msgchan chan string) {
	this.MsgChan = msgchan
}

func (this *Collect) SetSender(send types.ISend) {
	this.sender = send
}

func (this *Collect) send(msg string) {
	if len(msg) == 0 {
		return
	}
	for {
		if _, errNum, err := this.sender.Send(msg); err == nil {
			atomic.AddInt32(&this.SendSuccesCount, 1)
			if errNum != 0 {
				atomic.AddInt32(&this.ErrCount, int32(errNum))
			}
			break
		} else {
			atomic.AddInt32(&this.SendFailCount, 1)
			time.Sleep(3 * time.Second)
			this.ctx.Log.Error("发送失败", zap.Error(err))
			this.ErrMsg.Store(fmt.Sprintf("%s", err))
		}
	}
}

func (this *Collect) Run() {
	go this.Clock()
	for {
		//时间到了 发送
		//数量够了 发送
		//发送异常 持续自旋
		select {
		case <-this.act:
			//时间到了进行发送
			//发送会阻塞

			this.send(this.MsgTxt.String())
			this.MsgTxt.Reset()
			this.MsgCount = 0

		case data := <-this.MsgChan:
			if this.ctx.Config.Collect.MaxCount == 0 {
				this.send(data)
			} else {
				if this.MsgCount != 0 {
					this.MsgTxt.WriteString("\n")
				}
				this.MsgTxt.WriteString(data)
				this.MsgCount += 1

				if this.MsgCount == this.ctx.Config.Collect.MaxCount {
					this.send(this.MsgTxt.String())
					this.MsgTxt.Reset()
					this.MsgCount = 0
				}

			}
		}
	}
}

func (this *Collect) Clock() {
	for {
		select {
		case <-time.After(time.Second):
			select {
			case this.act <- struct{}{}:
			default:
			}
		}
	}
}

func (this *Collect) GetStatus() *types.Status {
	return &types.Status{
		SendSuccesCount: atomic.LoadInt32(&this.SendSuccesCount),
		SendFailCount:   atomic.LoadInt32(&this.SendFailCount),
		ErrCount:        atomic.LoadInt32(&this.ErrCount),
		StartTime:       this.startTime,
		LongTime:        time.Now().Sub(this.startTime).String(),
	}
}
