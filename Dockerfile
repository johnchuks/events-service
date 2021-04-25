ARG GO_VERSION=1.15

FROM golang:${GO_VERSION}-alpine

ENV GO111MODULE=on

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /app

# install compile daemon for hotreload
RUN go get github.com/githubnemo/CompileDaemon
