FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go build -o /bot ./cmd/bot

CMD ["/bot"]
