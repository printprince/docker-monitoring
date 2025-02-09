FROM golang:1.23-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем модули
COPY go.mod go.sum ./
RUN go mod download

# Копируем код
COPY . .

# Сборка приложения
RUN go build -o main ./cmd/main.go

# Используем минимальный образ Alpine для финального контейнера
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /root/

# Добавляем корневые сертификаты (на случай HTTPS-запросов)
RUN apk add --no-cache ca-certificates

# Копируем собранный бинарник из builder-слоя
COPY --from=builder /app/main .

# Открываем порт
EXPOSE 8082

# Запускаем приложение
CMD ["./main"]
