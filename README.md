# mqtt

盘古框架`Mqtt`集成，方便用户在`盘古`框架中接入`Mqtt`功能

> 本框架对`Mqtt`做了一定的封装，使用参看使用文档

## 快速开始

`Pangu`使用非常简单，只需要定义两部分代码`启动方法`和`启动器`以及`依赖关系`，对于`依赖关系`

- 可以没有依赖，如果没有，就不需要处理依赖
- 一般稍微大型一点的系统都会有复杂的依赖关系，而这正是`盘古`框架的强项
- 对于没有依赖或者依赖非常少的项目，使用`盘古`反而会使代码变多，这个需要大家酌情考虑

`启动方法`代码如下

```go
package main

import (
    `github.com/pangum/pangu`
    `github.com/pangum/pangum`
)

func main() {
    panic(pangu.New(
        pangu.Named(`ziyunix`),
        pangu.Banner(`Ziyunix Server`, pangu.BannerTypeAscii),
    ).Run(newBootstrap))
}
```

`启动器`的代码如下

```go
package main

type bootstrap struct {
    application *pangu.Application
}

func newBootstrap(application *pangu.Application) pangu.Bootstrap {
    return &bootstrap{
        application: application,
    }
}
```

`依赖关系`建议像如下处理

```go
package rest

import (
    `github.com/pangum/pangu`
)

type Server struct{}

func newServer( /* 如果有依赖，可以在这里增加依赖：api *Api */) *Server {
    return new(Server)
}

func init() {
    pangu.New().Musts(
        newServer,
        // 其它依赖
        // ...
    )
}
```

> `Pangu`有非常多的配置项，请参看[**使用文档**](https://pangu.pangum.tech)

## 文档

[点击这里查看最新文档](https://pangu.pangum.tech)

## 使用示例

[点击这里查看最新的代码示例](example)

## 交流

![微信群](doc/.vuepress/public/communication/wxwork.jpg)

## 捐助

![支持宝](https://github.com/storezhang/donate/raw/master/alipay-small.jpg)
![微信](https://github.com/storezhang/donate/raw/master/weipay-small.jpg)

