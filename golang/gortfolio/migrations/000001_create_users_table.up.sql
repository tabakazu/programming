CREATE TABLE users (
  id int NOT NULL AUTO_INCREMENT,
  email varchar(255) UNIQUE NOT NULL,
  created_at datetime,
  updated_at datetime,
  deleted_at datetime,
  PRIMARY KEY (id)
);
