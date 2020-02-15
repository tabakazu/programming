# Golang API Server
MEMO:
Create DB
```
$ mysql -u root -e 'CREATE DATABASE golang_api_srv_dev;'
```

Create new migrate file
```
$ migrate create -ext sql -dir db/migrations -seq create_posts
```

Edit migration file & Migrate
```
$ vi db/migrations/000001_create_posts.up.sql
$ migrate -database 'mysql://root:@/golang_api_srv_dev' -path ./db/migrations/ goto 000001
```

Insert sample data with SQL
```sql
INSERT INTO posts(title, body) VALUES ("first post title", "first!!!");
```

Run sandbox app
```
$ go run sandobox.go
```

Run main app
```
$ go run main.go
```
