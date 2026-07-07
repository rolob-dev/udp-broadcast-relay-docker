# Builder
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -o ssdp-relay \
    ./cmd/ssdp-relay

# Runtime
FROM alpine:3.22

WORKDIR /

COPY --from=builder /app/ssdp-relay .

ENTRYPOINT ["./ssdp-relay"]