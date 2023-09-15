FROM golang:1.19

WORKDIR /app

RUN go mod init github.com/curso-fullcycle-3-0/23-continuos-integration

COPY . .

RUN GOOS=linux go build ./cmd/main.go

CMD ["./main"]

