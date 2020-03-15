CREATE TABLE posts (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  created_at datetime,
  updated_at datetime,
  deleted_at datetime,
  PRIMARY KEY (id)
);
