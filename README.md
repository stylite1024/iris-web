# iris-web

[中文](./README-zh.md)

iris-web is a web template based on [vue3](https://vuejs.org) and [iris](https://github.com/kataras/iris) in  [Go](https://go.dev/). It does not require packaging of front-end static files, HTML directly uses Vue3 axios、element-plus ui， Embed into go binary file.

## Project Layout
See: [project-layout](https://github.com/golang-standards/project-layout)
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

## Getting started

### Prerequisites

Requires [Go](https://go.dev/) version [1.22.0](https://go.dev/doc/devel/release#go1.22.0) or above.

### Develop

```sh
git clone https://github.com/stylite1024/iris-web.git
cd iris-web
go mod tidy
go run main.go
```
Then visit [`0.0.0.0:8080`](http://0.0.0.0:8080) in your browser to see the response!

### Deploy

```sh
# build binary file
make build

# build docker image
make build-image
```

## Thanks
- [go](https://github.com/golang/go)
- [iris](https://github.com/kataras/iris)
- [vuejs](https://github.com/vuejs/vue)
- [axios](https://github.com/axios/axios)
- [element-plus-ui](https://github.com/element-plus/element-plus)
- [cobra](https://github.com/spf13/cobra)
- [cobra-cli](https://github.com/spf13/cobra-cli)
- [viper](https://github.com/spf13/viper)

## Issues

Found an error? Is there something meaningless? Initiate an [issue](https://github.com/stylite1024/iris-web/issues) to me, thank you!

## LICENSE

[MIT](https://github.com/stylite1024/iris-web/blob/main/LICENSE)
