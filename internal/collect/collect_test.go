package collect

import (
	"errors"
	"fmt"
	"log"
	"observelog/internal/collect/collectcfg"
	"observelog/internal/config"
	"observelog/internal/context"
	"sync"
	"testing"
	"time"
)

type Sender struct {
	Status      int
	DelaySecond time.Duration
	sync.RWMutex
}

func (this *Sender) SetStatus(s int) {
	this.Lock()
	defer this.Unlock()
	this.Status = s
}

func (this *Sender) SetDelay(s int) {
	this.Lock()
	defer this.Unlock()
	this.DelaySecond = time.Duration(s)
}

func (this Sender) Send(msg string) (int, int, error) {
	this.RLock()
	defer this.RUnlock()
	log.Println("发送数据\n", msg)
	time.Sleep(time.Second * this.DelaySecond)
	if this.Status == 0 {
		return 0, 0, errors.New("this is request error")
	} else {
		return 0, 0, nil
	}
}

func TestCollect(t *testing.T) {
	ctx := &context.AppCtx{}
	ctx.Config = config.Config{
		Collect: collectcfg.Cfg{
			MaxCount: 5,
		},
	}

	var msgchan = make(chan string)
	var sender = &Sender{}

	a := NewCollect(ctx)
	a.SetMsgChan(msgchan)
	a.SetSender(sender)
	go a.Run()

	go func() {
		ct := 0
		for {
			select {
			case msgchan <- fmt.Sprintf(`{"count":%d,"name":"mock"}`, ct):
				log.Println("采集数据:", fmt.Sprintf(`{"count":%d,"name":"mock"}`, ct))
			default:
				log.Println("采集数据阻塞")
			}
			ct += 1
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			log.Println(a.GetStatus())
			time.Sleep(time.Second)
		}
	}()

	sender.SetDelay(0)
	sender.SetStatus(1)

	time.Sleep(time.Second * 20)

}
