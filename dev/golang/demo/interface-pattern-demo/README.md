init sql
```
CREATE DATABASE IF NOT EXISTS simple_rest_api_dev;

USE simple_rest_api_dev;

SHOW COLUMNS FROM posts;

INSERT INTO posts (title, body) VALUES ("Test Title", "Test Body");

SELECT * FROM posts;
```

run app
```
go run main.go
```
