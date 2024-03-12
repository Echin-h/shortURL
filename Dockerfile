# 第一阶段：使用Go语言环境构建应用程序
FROM golang:1.21 as builder
ENV GOPROXY https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

# 将当前目录的内容复制到容器内的/app目录
COPY . .

# 构建应用程序，生成可执行文件
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main .

# 第二阶段：从一个空镜像创建最终的镜像
FROM alpine:3.12
WORKDIR /app
# 复制第一阶段构建的可执行文件到最终镜像的/app目录
COPY --from=builder /app/main /app/main
# 暴露应用程序监听的端口
EXPOSE 8080

# 设置容器启动时运行的命令
CMD ["/app/main"]