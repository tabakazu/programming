# Rails (Ruby + Node.js) Multi-stage Build Dockerfile
- Ruby 2.3.7
- Rails 5.2
- Node.js 8.11.3

## Build
```
$ docker build -t rails5 .
```

## Rails new
```
$ docker run -i -t rails5 bundle exec rails new . -d postgresql
```
