### Stage 1

FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o myapp server.go

### Stage 2 
FROM debian:bookworm-slim

WORKDIR /app

COPY .env .env

COPY --from=builder /app/myapp .

EXPOSE 8000

CMD ["./myapp"]