FROM golang:alpine

MAINTAINER alexbosworth

ENV GIN_MODE=release

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build

EXPOSE 8080

ENTRYPOINT ["./grpc-proxy"]
