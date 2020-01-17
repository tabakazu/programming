Install
```
$ brew install golang-migrate
$ which migrate
/usr/local/bin/migrate
```

Create DB
```
$ mysql -u root -e 'CREATE DATABASE golang_migrate_demo;'
```

Create new migrate file
```
$ migrate create -ext sql -dir db/migrations -seq create_posts_table
# $ migrate create -ext sql -dir db/migrations -seq alter_posts_table
```

Edit migrate file
```
$ vi db/migrations/000001_create_posts_table.up.sql

CREATE TABLE posts (
  id int NOT NULL,
  title varchar(255) NOT NULL,
  body text NOT NULL,
  created_at datetime,
  updated_at datetime,
  deleted_at datetime,
  PRIMARY KEY (id)
);

$ vi db/migrations/000001_create_posts_table.down.sql

DROP TABLE posts;
```

Run migration
```
$ migrate -database 'mysql://root:@/golang_migrate_demo' -path ./db/migrations up
# $ migrate -database 'mysql://root:@/golang_migrate_demo' -path ./db/migrations down
# $ migrate -database 'mysql://root:@/golang_migrate_demo' -path ./db/migrations/ goto 000001
```

Check migrate status
```
$ migrate -database 'mysql://root:@/golang_migrate_demo' -path ./db/migrations version
```
