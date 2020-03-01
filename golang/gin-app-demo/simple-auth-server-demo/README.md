# Simple Auth Server

## Create DB
```
$ mysql -u root -e 'CREATE DATABASE go_simple_auth_server_dev;'
```

# Create migration file
```
$ migrate create -ext sql -dir db/migrations -seq create_users_table
```

# Migrate schema
```
$ migrate -database 'mysql://root:@/go_simple_auth_server_dev' -path ./db/migrations up
```

## API
### Register
```bash
$ curl -X POST http://localhost:8080/register \
       -H 'content-type: application/json' \
       -d '{ "email": "taba@test.com", "password": "passw0rd" }'
```

### Authenticate
```bash
$ curl -X POST http://localhost:8080/authenticate \
       -H 'content-type: application/json' \
       -d '{ "email": "taba@test.com", "password": "passw0rd" }'
```
