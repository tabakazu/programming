FROM golang:1.12

ENV GO111MODULE=on

WORKDIR /go/src/dockerize-goapp-demo

COPY . .

RUN go get github.com/pilu/fresh

CMD ["fresh"]

EXPOSE 8080