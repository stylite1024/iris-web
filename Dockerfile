ARG GOVERSION="1.22.0"

# 构建镜像
FROM golang:${GOVERSION}-alpine AS Builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 工作目录：/build
WORKDIR /build

# 将代码复制到容器中的/build下
COPY . .

# 将我们的代码编译成二进制可执行文件app
RUN go mod download && go build -ldflags "-s -w" -a -installsuffix cgo -o app .

# 运行镜像
FROM scratch

# 将二进制文件从 /build 目录复制到这里
COPY --from=Builder /build/app /app

# 数据卷
VOLUME ["/data"]

# 声明服务端口
EXPOSE 8080

# 启动容器时运行的命令
CMD ["/app"]