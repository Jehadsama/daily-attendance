FROM golang:alpine AS builder
LABEL maintainer="Jehadsama<339364351@qq.com>"

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct \
    DAILY_ENV=dev

# 移动到工作目录：/src-app
WORKDIR /src-app

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件app
RUN go build cmd/main.go

###################
# 接下来创建一个小镜像
###################
FROM scratch

ENV DAILY_ENV=dev

WORKDIR /srcgo
COPY . ./

# 从builder镜像中把/src-app/main 拷贝到当前目录
COPY --from=builder /src-app/main /srcgo

# 需要运行的命令
ENTRYPOINT ["/srcgo/main", "cron"]