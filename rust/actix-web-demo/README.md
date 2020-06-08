# Actix web Demo

## Use Docker
```bash
$ docker-compose build
$ docker-compose run --rm --service-ports app bash
```

## Create a new project
```bash
docker > cargo init
```

## Setup hyper
Cargo.toml
```toml
[dependencies]
actix-web = "2"
actix-rt = "1"
```

## Run Server
```bash
docker > cargo run
```

http://localhost:3000/1/name/index.html
