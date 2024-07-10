Web chat with WebSocket and notifications for logged-in and logged-out users and using emojis in Go and JavaScript. 
<br />
Now, if you want to use a chat with messages and track connected and disconnected users using Redis, check out this <a href="https://github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang">project</a>
<br />
1 - Run: cmd/main.go<br />
2 - access via browser: http://localhost:8080/chat<br />
2 - being able to open in multiple tabs and connect multiple users<br />
3 - connect user and send message

<br/>

You can also run it through the dockerfile:<br />

 ```
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

 ```