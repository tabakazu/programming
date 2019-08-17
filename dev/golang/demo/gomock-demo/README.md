```
$ go get github.com/golang/mock/gomock
$ go get github.com/golang/mock/mockgen
$ mkdir -p mocks
```

```
$ mockgen -destination=mocks/mock_post.go -package=mocks github.com/tabakazu/gomock-demo/usecase Post
$ go test -v github.com/tabakazu/gomock-demo/controller
```