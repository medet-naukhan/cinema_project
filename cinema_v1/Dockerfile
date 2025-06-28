FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o cinema-app ./cmd/main.go

FROM alpine:latest
COPY --from=builder /app/cinema-app .
CMD ["./cinema-app"]