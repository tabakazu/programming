FROM golang:1.11-alpine
COPY ./api /go/src/api/
WORKDIR /go/src/api/
RUN apk add --no-cache git && \
  go get -u github.com/codegangsta/gin && \
  go get -u github.com/golang/dep/cmd/dep
