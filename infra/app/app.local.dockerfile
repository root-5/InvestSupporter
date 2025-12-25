FROM golang:alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install golang.org/x/tools/cmd/godoc@latest

# ボリュームマウントはRUNの直前で行われる

CMD ["sh", "-c", "godoc -http=:8081 & air -c .air.toml"]
