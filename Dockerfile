# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go mod verify

COPY index.html ./
COPY *.go ./

RUN go build -o /calc-server

EXPOSE 80

CMD ["/calc-server"]