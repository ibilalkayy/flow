FROM golang:1.22.0-alpine3.19

ARG SERVICE_NAME

RUN mkdir -p /app

WORKDIR /app

COPY cmd/transaction/main/main.go /app
COPY go.mod /app
COPY go.sum /app

RUN go mod download

COPY . /app

RUN go build -o flowt .

CMD [ "./flowt" ]