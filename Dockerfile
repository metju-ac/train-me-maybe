FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

WORKDIR /app/cmd/multi-user
RUN go build -o /train-me-maybe-app

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /train-me-maybe-app .
COPY config.toml .

CMD ["./train-me-maybe-app"]