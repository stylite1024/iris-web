# iris-web

[English](./README.md)

## 简介

iris-web是一个基于go iris和vue3框架的web模板。它采用的是前后端不分离模式，可用于快速开发web应用。

## 目录结构

```
.
├── cmd
│   ├── root.go
│   └── version.go
├── config
│   └── info.go
├── docker-compose.yml
├── Dockerfile
├── docs
├── go.mod
├── go.sum
├── internal
│   ├── handler
│   │   └── handler.go
│   ├── logger
│   │   └── logger.go
│   ├── logic
│   └── router
│       └── router.go
├── LICENSE
├── main.go
├── Makefile
├── pkg
│   └── util
│       └── response.go
├── README.md
├── README-zh.md
├── scripts
│   └── docker-entrypoint.sh
└── web
    ├── static
    │   ├── favicon.ico
    │   └── index.html
    └── static.go
```

## 快速开始

### 要求

golang版本1.22.0及以上

### 开发

```sh
git clone https://github.com/stylite1024/iris-web.git
cd iris-web
go mod tidy
go run main.go
```

浏览器访问：http://0.0.0.0:8080

### 部署

```sh
# build binary file
make build

# build docker image
make build-image
```

## 致谢

- [go](https://github.com/golang/go)
- [iris](https://github.com/kataras/iris)
- [vuejs](https://github.com/vuejs/vue)
- [axios](https://github.com/axios/axios)
- [element-plus-ui](https://github.com/element-plus/element-plus)
- [cobra](https://github.com/spf13/cobra)
- [cobra-cli](https://github.com/spf13/cobra-cli)
- [viper](https://github.com/spf13/viper)

## 提问

Found an error? Is there something meaningless? Initiate an [issue]() to me, thank you!

## 开源协议

[MIT]()
