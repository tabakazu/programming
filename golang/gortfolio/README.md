## Setup
```bash
$ brew cask install docker

$ docker-compose build
$ docker-compose up -d

$ docker-compose run --rm migrate -database 'mysql://root:root@tcp(mysql:3306)/gortfolio_dev' -path ./migrations up

$ docker-compose exec goapp bash
$ go run main.go
# http://localhost:8090/ping
```

## Migration
```
$ docker-compose run --rm migrate create -ext sql -dir migrations -seq create_users_table
$ docker-compose run --rm migrate -database 'mysql://root:root@tcp(mysql:3306)/gortfolio_dev' -path ./migrations up
$ docker-compose run --rm migrate -database 'mysql://root:root@tcp(mysql:3306)/gortfolio_dev' -path ./migrations down
```

## Down
```
$ docker-compose down --rmi all --volumes
```
