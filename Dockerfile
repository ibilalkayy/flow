FROM golang:1.22.0-alpine3.19

RUN mkdir -p /app
WORKDIR /app

COPY go.mod /app
COPY go.sum /app

COPY . /app

RUN go mod download

RUN go build -o flow .

CMD ["./flow"]