# Используем официальный образ Golang
FROM golang:latest

# Установка рабочей директории
WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

# Копирование исходного кода в контейнер
COPY . .


RUN go build -o main cmd/main.go

# Указание порта, который будет прослушиваться приложением
EXPOSE 8090

CMD ["./main"]
