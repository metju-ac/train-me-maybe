# Step 1: Build the Backend
FROM golang:1.23.4-alpine AS backend-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

WORKDIR /app/cmd/multi-user
RUN go build -o /train-me-maybe-app

# Step 2: Build the Frontend
FROM node:20-alpine AS frontend-builder

WORKDIR /frontend

COPY ./frontend/package.json ./
RUN npm install

COPY ./frontend ./
RUN npm run build

# Step 3: Create the Final Image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=backend-builder /train-me-maybe-app .
COPY config.toml .
COPY migrations ./migrations

COPY --from=frontend-builder /frontend/dist ./static

CMD ["./train-me-maybe-app"]