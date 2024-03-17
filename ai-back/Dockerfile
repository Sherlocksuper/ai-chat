# 第一阶段：构建Go可执行文件
FROM golang:alpine AS builder
LABEL stage=gobuilder

# 禁用CGO，确保构建过程独立于宿主机
ENV CGO_ENABLED=0
ENV GOOS=linux

# 设置代理（可选）
ENV GOPROXY=https://goproxy.cn

# 设置工作目录
WORKDIR /build/your_project

# 复制go.mod和go.sum并下载依赖
COPY go.mod .
COPY go.sum .
RUN go mod download

# 复制整个项目并构建可执行文件
COPY . .
RUN go build -ldflags="-s -w" -o /app/main


# 第二阶段：构建最终镜像
FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ=Asia/Shanghai

# 设置工作目录
WORKDIR /app

# 复制可执行文件
COPY --from=builder /app/main /app/main

# 启动服务
CMD ["/app/main"]
