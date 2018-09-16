# Gin-gonic (Golang)
- Go 1.11.0

## Build
```
$ docker build -t ginapp .
```

## Run Ginapp
```
$ docker run --rm -it -p 3001:3001 ginapp sh
/go/src/api # dep init
/go/src/api # dep ensure
/go/src/api # gin -i run
```
http://localhost:3001/ping
