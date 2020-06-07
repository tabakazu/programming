# hyper Demo

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
hyper = "0.13"
tokio = { version = "0.2", features = ["full"] }
```

## Run Server
```bash
docker > cargo run
```
