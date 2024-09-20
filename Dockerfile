# Используем официальный образ Go
FROM golang:1.23-alpine

# Определяем порт, на котором работает приложение
EXPOSE 8080

# Создаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь исходный код в контейнер
COPY . .

# Собираем бинарный файл
RUN go build -o main ./cmd/tenderService

# Запускаем приложение
CMD ["./main"]