FROM golang:1.22.0-alpine3.19 as builder

WORKDIR /app

COPY go.mod /app
COPY go.sum /app

RUN go mod download

COPY . /app

RUN go build -ldflags="-w -s" -o flow

FROM busybox

COPY --from=builder /app/flow /flow

CMD ["sh", "-c", "./flow && while true; do echo 'App is running'; sleep 10; done"]
