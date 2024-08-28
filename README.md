Web chat with WebSocket and notifications for logged-in and logged-out users and using emojis in Go and JavaScript. 
<br /><br/>
Now, if you want to use a chat with messages and track connected and disconnected users using Redis, check out this <a href="https://github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang">project</a>
<br /><br/>
1 - Run: cmd/main.go<br />
2 - access via browser: http://localhost:8080/chat<br />
2 - being able to open in multiple tabs and connect multiple users<br />
3 - connect user and send message

<br/>

You can also run it through the dockerfile:<br />

 ```
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .

RUN go mod download
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o main ./cmd/main.go

FROM scratch
WORKDIR /app


COPY --from=builder /app/main /app/
COPY --from=builder /app/cmd/.env /app/

CMD ["./main"]
 ```