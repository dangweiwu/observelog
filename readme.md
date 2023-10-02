<h1 align="center">Observe Log</h1>
<p align="center">
<img alt="Static Badge" src="https://img.shields.io/badge/Go- 1.9-blue">
<img alt="Static Badge" src="https://img.shields.io/badge/Gin- 1.8-blue">

<img alt="Static Badge" src="https://img.shields.io/badge/OpenObserve- 0.60 -gren">
<img alt="Static Badge" src="https://img.shields.io/badge/issue- open -gren">
<img alt="Static Badge" src="https://img.shields.io/badge/license- MIT-blue">
<h4 align="center">简单、稳定的日志采集工具，OpenPbserve专用</h4>
</p>

---

### 说明：

本项目为[OpenObserve](https://openobserve.ai/)专用日志采集工具，具有使用简单、稳定、资源使用低等特点。

使用本工具前请先安装[OpenObserve](https://openobserve.ai/)服务。
### 安装:

```go install github.com/dangweiwu/```

### 配置文件
config/logconfig.yaml
``` yaml
Host: localhost:8102 //日志监控服务端口，可查看工作状态
user: root           //日志监控服务账号
password: root       //日志监控服务密码
Logx:
  LogName: ./collect.log     //日志监控服务的日志
  Level: debug               //日志级别
  OutType: all               //日志输出配置 console,file,all
  Formatter: txt             //日志格式 txt,json

Observe:
  Host: http://localhost:5080  //OpenObserve服务地址
  UserName: admin@qq.com        //OpenObserve 账号
  Password: '123456'            //OpenObserve 密码
  Org: default                  //OpenObserve 组织
  Stream: server                //OpenObserve 流名称

Collect:
  MaxCount: 20                  //采集缓冲数量
  Path: ./demo.log              //采集日志地址
```

### 使用
```
observelog -f config.yaml 
```
该命令即可启动日志监控

*如果有多个日志需要监控请使用不同配置文件名以便区分*

### 采集状态

访问配置文件中的Host即可获取当前日志采集状态包括，发送数量，发送成功数量。


### 其他说明

- 所采集日志必须为json，非json日志将会被过滤掉
- MaxCount参数指定最大缓冲日志条数，当1s内，达到最大日志采集数量，会马上发送缓存日志，否则，1s时间达到，也会将日志进行发送。
- 本项目经过长期实践检验，每秒上千发送量也能长期稳定运行。



---
### License
© Dangweiwu, 2023~time.Now

Released under the MIT License
