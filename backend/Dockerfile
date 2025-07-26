FROM golang:1.24.4-alpine

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY ./src/go.mod ./src/go.sum ./

RUN go mod download

CMD ["air"]
