FROM golang:alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install golang.org/x/tools/cmd/godoc@latest

# ボリュームマウントはRUNの直前で行われる

CMD ["air", "-c", ".air.toml", "&&", "godoc", "-http=:8080"]