FROM golang:1.23-alpine AS builder

WORKDIR /app

# Копируем файлы модуля и скачиваем зависимости
COPY go.mod ./
RUN go mod download

# Копируем весь исходный код и собираем приложение
COPY . .
RUN go build -o pinger ./cmd/main.go

# Финальный образ
FROM alpine:latest
WORKDIR /root/

# Копируем только собранный бинарный файл
COPY --from=builder /app/pinger .
EXPOSE 8081

CMD ["./pinger"]
