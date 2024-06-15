FROM golang:alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

# ボリュームマウントはRUNの直前で行われる

CMD ["air", "-c", ".air.toml"]