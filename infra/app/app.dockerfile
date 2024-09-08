FROM golang:alpine AS builder

WORKDIR /app

# 依存関係をコピーし、ダウンロードする
COPY app/go.mod .
COPY app/go.sum .
RUN go mod download

# アプリケーションのソースをコピーする
COPY app .

# アプリケーションを main という名前でビルド
RUN go build -o main .

# 実行用のステージ
FROM alpine:latest

WORKDIR /root/

# ビルドしたバイナリをコピー
COPY --from=builder /app/main .

# index.html をコピー
COPY --from=builder /app/controller/api/index.html ./controller/api/index.html

# アプリケーションの起動コマンド実行
CMD ["./main"]