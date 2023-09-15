FROM golang:1.19

WORKDIR /app

RUN go mod init teste

COPY . .

RUN GOOS=linux go build ./cmd/main.go

CMD ["./main"]

