FROM golang:1.23.1 AS builder
ARG TARGETOS=linux
ARG TARGETARCH=amd64
# 添加新的构建参数用于指定二进制文件名
ARG BINARY_NAME=query-wizard

# 设置 Go 模块代理为官方代理
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /workspace
ENV GO111MODULE=auto

# 禁用CGO并启用静态编译
ENV CGO_ENABLED=0

COPY . .

# 整理模块依赖
RUN go mod tidy

# 使用构建参数 BINARY_NAME 来指定输出的二进制文件名
RUN CGO_ENABLED=0 go build -a -o ${BINARY_NAME} main.go

# 调试：列出生成的文件及其权限
RUN ls -l /workspace

FROM alpine:3.1
ARG BINARY_NAME=query-wizard

USER 0
WORKDIR /

# 从构建阶段复制生成的静态二进制文件，使用相同的 BINARY_NAME
COPY --from=builder /workspace/${BINARY_NAME} .
COPY --from=builder /workspace/config/ config/
COPY --from=builder /workspace/handlers/ handlers/
COPY --from=builder /workspace/influxdb/ influxdb/
COPY --from=builder /workspace/k8s/ k8s/

# 确保二进制文件具有执行权限
RUN chmod +x /${BINARY_NAME} && \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 使用构建参数指定的二进制文件作为入口点
#ENTRYPOINT ["/${BINARY_NAME}"]
ENTRYPOINT ["/query-wizard"]


