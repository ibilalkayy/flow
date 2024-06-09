FROM golang:1.22.0-alpine3.19

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

COPY . /app

RUN go mod download

RUN go build -o flow .

CMD ["sh", "-c", "./flow && while true; do echo 'App is running'; sleep 10; done"]
