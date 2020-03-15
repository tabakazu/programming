#　ドメインモデリング

## データベースセットアップ
### データベースの作成
```bash
$ mysql -u root -e 'CREATE DATABASE domain_modeling_demo;'
```

### Migration ファイルの作成
```bash
$ migrate create -ext sql -dir migrations -seq create_posts_table
```

### Migration ファイルの編集 
migrations/000001_create_posts_table.up.sql
```sql
CREATE TABLE posts (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  created_at datetime,
  updated_at datetime,
  deleted_at datetime,
  PRIMARY KEY (id)
);
```

migrations/000001_create_posts_table.down.sql
```sql
DROP TABLE posts;
```

### データベーススキーマの更新

```bash
$ migrate -database 'mysql://root:@/domain_modeling_demo' -path ./migrations up
```
