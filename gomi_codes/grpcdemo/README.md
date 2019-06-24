```
$ brew install protobuf

$ go get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/proto
$ go get -u github.com/golang/protobuf/protoc-gen-go

$ protoc -I pb/ pb/hello.proto --go_out=plugins=grpc:pb

$ go run server.go
$ go run client.go
```
