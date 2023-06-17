package option

import (
	"github.com/gin-gonic/gin"
	"gs/pkg/yamconfig"
	"observelog/internal/collect"
	"observelog/internal/config"
	"observelog/internal/context"
)

type RunServe struct {
}

func (this *RunServe) Execute(args []string) error {
	//配置参数
	var c config.Config
	yamconfig.MustLoad(Opt.ConfigPath, &c)

	//资源初始化
	sc, err := context.NewAppCtx(c)
	if err != nil {
		panic(err)
	}

	//启动日志采集
	clt := collect.NewCollect(sc)
	til := collect.NewTailx(sc)
	req := collect.NewReqObserve(sc)
	clt.SetMsgChan(til.GetMsgChan())
	clt.SetSender(req)
	go clt.Run()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.BasicAuth(gin.Accounts{c.User: c.Password}))
	router.GET("/status", func(c *gin.Context) {
		status := clt.GetStatus()
		c.JSON(200, status)
	})

	//启动
	router.Run(c.Host)
	return nil
}
