CREATE TABLE customers (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(255) NOT NULL,
  street varchar(255) NOT NULL,
  city varchar(255) NOT NULL,
  country varchar(255) NOT NULL,
  created_at datetime,
  updated_at datetime,
  deleted_at datetime,
  PRIMARY KEY (id)
);
