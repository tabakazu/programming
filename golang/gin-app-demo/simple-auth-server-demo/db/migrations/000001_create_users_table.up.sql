CREATE TABLE users (
  id int UNIQUE NOT NULL,
  email varchar(255) UNIQUE NOT NULL,
  encrypted_password varchar(255) NOT NULL,
  api_token varchar(255) UNIQUE NOT NULL,
  created_at datetime,
  updated_at datetime,
  deleted_at datetime,
  PRIMARY KEY (id)
);
