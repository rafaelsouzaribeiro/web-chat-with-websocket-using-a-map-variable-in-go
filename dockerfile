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