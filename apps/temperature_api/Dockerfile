# Используем официальный образ Go в качестве базового
FROM golang:1.22-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod для кэширования зависимостей
COPY go.mod ./
RUN go mod download

# Копируем остальные файлы приложения
COPY . .

# Компилируем приложение
RUN go build -o temperature-api .

# Используем минимальный образ для запуска
FROM alpine:latest

# Устанавливаем необходимые зависимости для работы
RUN apk --no-cache add ca-certificates

# Копируем скомпилированное приложение из образа builder
COPY --from=builder /app/temperature-api .

EXPOSE 8082

# Указываем команду для запуска приложения
CMD ["./temperature-api"]