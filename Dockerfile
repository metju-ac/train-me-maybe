ARG GOLANG_VERSION="1.23.4-alpine"
ARG NODE_VERSION="22-alpine"

# Step 1: Build the Backend
FROM golang:${GOLANG_VERSION} AS backend-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download -x

COPY . ./

WORKDIR /app/cmd/multi-user
RUN go build -mod=readonly -v -o /train-me-maybe-app
 
# Step 2: Build the Frontend
ARG NODE_VERSION
FROM node:${NODE_VERSION} AS frontend-builder

WORKDIR /frontend

COPY ./frontend/package*.json ./
RUN --mount=type=cache,target=~/.npm npm ci

COPY ./frontend ./
RUN npm run build

# Step 3: Create the Final Image
FROM alpine:latest
RUN --mount=type=cache,target=/var/cache/apk,sharing=locked apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=backend-builder /train-me-maybe-app .
COPY config.toml .
COPY migrations ./migrations

COPY --from=frontend-builder /frontend/dist ./static

CMD ["./train-me-maybe-app"]