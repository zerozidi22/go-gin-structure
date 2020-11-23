FROM golang:1.15 AS builder
WORKDIR /builder

ENV CGO_ENABLED=0

COPY go.mod ./
RUN go mod download

COPY . ./
RUN go build .

EXPOSE 8000
ENTRYPOINT= ["/bin/test"]