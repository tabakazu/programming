```bash
# start
$ docker-compose up -d

# use MySQL
$ docker-compose exec mysql bash
$ mysql --user root --host 127.0.0.1 --port=13306 < tmp/data/init.sql
$ mysql --user root --host 127.0.0.1 --port=13306

# down
$ docker-compose down
```
