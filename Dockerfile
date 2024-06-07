FROM golang:alpine
LABEL maintainer="Jehadsama<339364351@qq.com>"

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct \
    DAILY_ENV=prod

# 移动到工作目录：/build
WORKDIR /src-app

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件app
RUN go build cmd/main.go

EXPOSE 9000

# 需要运行的命令
ENTRYPOINT ["/src-app/main", "cron"]