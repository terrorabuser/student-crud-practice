# syntax=docker/dockerfile:1

###########################
# 🔨  Stage 1 — build
###########################
FROM golang:1.24-alpine AS builder

# 1. Устанавливаем зависимости
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# 2. Копируем исходники и собираем бинарь
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o server ./cmd   # поправьте путь, если main.go лежит иначе

###########################
# 🚀  Stage 2 — runtime
###########################
###########################
# 🚀  Stage 2 — runtime
###########################
FROM alpine:3.20

WORKDIR /app

# Копируем бинарь
COPY --from=builder /app/server .
COPY migrations ./migrations

EXPOSE 8080
ENTRYPOINT ["./server"]

