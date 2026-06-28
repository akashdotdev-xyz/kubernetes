# ----------------------------
# Stage 1 : Build
# ----------------------------
FROM golang:1.25.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o order-service ./cmd/server

# ----------------------------
# Stage 2 : Runtime
# ----------------------------
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/order-service .

EXPOSE 8080

CMD ["./order-service"]