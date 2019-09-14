```
$ go get github.com/golang/mock/gomock
$ go get github.com/golang/mock/mockgen
$ mkdir -p mocks
```

```
$ mockgen -destination=mocks/mock_post_usecase.go -package=mocks github.com/tabakazu/gomock-demo/controller PostUsecase
$ mockgen -destination=mocks/mock_post_repository.go -package=mocks github.com/tabakazu/gomock-demo/usecase PostRepository

$ go test -v github.com/tabakazu/gomock-demo/controller
```