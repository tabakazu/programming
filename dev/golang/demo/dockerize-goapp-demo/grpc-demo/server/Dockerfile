FROM golang:1.12.9-alpine3.10

RUN apk add --update --no-cache ca-certificates git

WORKDIR /go/src/app
ENV GO111MODULE=on

COPY ./server/go.mod .
COPY ./server/go.sum .
RUN go mod download

COPY pb pb

COPY ./server .