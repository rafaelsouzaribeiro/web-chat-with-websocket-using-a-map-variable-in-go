FROM golang:1.22.0 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

ENV HOST_NAME=0.0.0.0
ENV WS_ENDPOINT=/ws
ENV PORT=8080

COPY --from=builder /app/main .

EXPOSE $PORT

CMD ["./main"]
