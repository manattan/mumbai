# ビルドステージ
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/echo-server ./cmd/echo-server

# 実行ステージ
FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /bin/echo-server /bin/echo-server
EXPOSE 8080 50051
CMD ["/bin/echo-server"]