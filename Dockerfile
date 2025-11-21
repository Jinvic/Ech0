# =================== 构建阶段 ===================
FROM golang:1.25.3-alpine AS builder

WORKDIR /src

# 安装 Node.js 和 pnpm
RUN apk add --no-cache nodejs npm bash
RUN npm install -g pnpm

# 设置时区（可选）
ENV TZ=Asia/Shanghai

# 复制 go mod 文件，先下载依赖（缓存友好）
COPY go.mod go.sum ./
RUN go mod download

# 复制整个源码
COPY . .

# 构建前端
RUN cd web && pnpm install --frozen-lockfile && pnpm build --mode production

# 编译 Go 二进制（已自动 embed /web/dist）
# 使用 CGO_ENABLED=0 以生成静态二进制（兼容 Alpine）
RUN CGO_ENABLED=0 GOOS=linux go build -tags netgo -ldflags="-w -s" -o ech0 ./main.go

# =================== 最终镜像 ===================
FROM alpine:latest

WORKDIR /app

# 安装 ca-certificates（便于 HTTPS 请求）
RUN apk --no-cache add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 创建数据和备份目录
RUN mkdir -p /app/data /app/backup

# 从 builder 阶段复制二进制
COPY --from=builder /src/ech0 /app/ech0

# 设置权限
RUN chmod +x /app/ech0

EXPOSE 6277 6278

ENTRYPOINT ["/app/ech0"]
CMD ["serve"]