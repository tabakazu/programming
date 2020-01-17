CREATE TABLE posts (
  id int NOT NULL,
  title varchar(255) NOT NULL,
  body text NOT NULL,
  created_at datetime,
  updated_at datetime,
  deleted_at datetime,
  PRIMARY KEY (id)
);
