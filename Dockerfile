# Этап сборки
FROM golang:1.25 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Собираем статический бинарь
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Финальный контейнер (без лишнего мусора)
FROM scratch
WORKDIR /
COPY --from=builder /app/server /

# Указываем порт (для наглядности, не обязателен)
EXPOSE 8080

# Запускаем
ENTRYPOINT ["/server"]
