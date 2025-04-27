ARG GOLANG_VERSION="1.23.4-alpine"

# Step 1: Build the Backend
FROM golang:${GOLANG_VERSION} AS backend-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download -x

COPY . ./

WORKDIR /app/cmd/multi-user
RUN go build -mod=readonly -v -o /train-me-maybe-app

# Step 2: Create the Final Image
FROM alpine:latest
RUN --mount=type=cache,target=/var/cache/apk,sharing=locked apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=backend-builder /train-me-maybe-app .
COPY config.toml .
COPY migrations ./migrations

CMD ["./train-me-maybe-app"]