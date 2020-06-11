# diesel Demo

## Use Docker
```bash
$ docker-compose build
$ docker-compose run --rm --service-ports app bash
```

## Create a new project
```bash
docker > cargo init --bin
docker > cargo install diesel_cli --no-default-features --features mysql
docker > diesel setup
```

## Setup
Cargo.toml
```toml
[dependencies]
diesel = { version = "1.4.4", features = ["mysql"] }
dotenv = "0.15.0"
```

## Migrate Schema
```bash
# create migrate file
docker > diesel migration generate create_posts
# apply
docker > diesel migration run
```

## Run Server
```bash
docker > cargo run
```
