FROM golang:1.21 as builder

WORKDIR /app
COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init || true

RUN go build -o server .

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
