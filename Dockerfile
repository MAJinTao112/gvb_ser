# ================================
# Go 应用多阶段构建（无语法声明版）
# ================================

# 允许自定义 Go 版本，默认 1.22，可在构建时用 --build-arg GO_VERSION=1.23 覆盖
ARG GO_VERSION=1.24.0

# ====== Build Stage ======
FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /app

# 安装构建依赖
RUN apk add --no-cache git ca-certificates tzdata

# 设置 Go 代理并预下载依赖
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod download

# 拷贝全部源码
COPY . .

# 构建二进制（静态编译，减小运行镜像依赖）
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -trimpath -ldflags="-s -w" -o /app/gvb_server ./main.go

# ====== Run Stage ======
FROM alpine:3.20

WORKDIR /app

# 安装运行依赖（证书/时区），并创建非 root 用户
RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S app && adduser -S app -G app \
    && mkdir -p /app/uploads/file \
    && chown -R app:app /app

# 拷贝构建产物和必要文件
COPY --from=builder /app/gvb_server /app/gvb_server
COPY settings.yaml /app/settings.yaml
COPY avatar.png /app/avatar.png
COPY uploads /app/uploads

ENV TZ=Asia/Shanghai

# 暴露端口（根据 settings.yaml 中 system.port 控制）
EXPOSE 8080

USER app

# 程序入口
ENTRYPOINT ["/app/gvb_server"]
