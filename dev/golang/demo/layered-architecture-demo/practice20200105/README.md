# 2020/01/05

Initial
```sql
CREATE DATABASE IF NOT EXISTS godev20200105;

```

Generate mock
```
$ mockgen -destination=mocks/mock_item_repository.go -package=mocks github.com/tabakazu/practice20200105/domain/repository ItemRepository
```
