# WeChat SDK for Go

![Go](https://github.com/rmrf/wechat/workflows/Go/badge.svg?branch=release-2.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/rmrf/wechat)](https://goreportcard.com/report/github.com/rmrf/wechat)
[![pkg](https://img.shields.io/badge/dev-reference-007d9c?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/rmrf/wechat/v2?tab=doc)
![version](https://img.shields.io/badge/version-v2-green)

使用 Golang 开发的微信 SDK，简单、易用。

> 注意：当前版本为 v2 版本，v1 版本已废弃

## 文档 && 例子

[API 列表](https://github.com/rmrf/wechat/tree/v2/doc/api)

[Wechat SDK 2.0 文档](https://silenceper.com/wechat)

[Wechat SDK 2.0 例子](https://github.com/gowechat/example)

## 快速开始

```
import "github.com/rmrf/wechat/v2"
```

以下是一个微信公众号处理消息接收以及回复的例子：

```go
//使用memcache保存access_token，也可选择redis或自定义cache
wc := wechat.NewWechat()
memory := cache.NewMemory()
cfg := &offConfig.Config{
    AppID:     "xxx",
    AppSecret: "xxx",
    Token:     "xxx",
    //EncodingAESKey: "xxxx",
    Cache: memory,
}
officialAccount := wc.GetOfficialAccount(cfg)

// 传入request和responseWriter
server := officialAccount.GetServer(req, rw)
//设置接收消息的处理方法
server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {

    //回复消息：演示回复用户发送的消息
    text := message.NewText(msg.Content)
    return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
})

//处理消息接收以及回复
err := server.Serve()
if err != nil {
    fmt.Println(err)
    return
}
//发送回复的消息
server.Send()

```

## 目录说明

- officialaccount: 微信公众号 API
- miniprogram: 小程序 API
- minigame:小游戏 API
- pay:微信支付 API
- openplatform:开放平台 API
- work:企业微信
- aispeech:智能对话
- doc: api 文档

## 贡献

- 在[API 列表](https://github.com/rmrf/wechat/tree/v2/doc/api)中查看哪些 API 未实现
- 提交 issue，描述需要贡献的内容
- 完成更改后，提交 PR

## 公众号

![img](https://silenceper.oss-cn-beijing.aliyuncs.com/qrcode/search_study_program.png)

## License

Apache License, Version 2.0
