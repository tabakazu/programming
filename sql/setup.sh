# for MySQL 5.7 container
mysql --user root --host 127.0.0.1 --port=13306 < seed/01_setup_tables__mysql_5_7.sql
mysql --user root --host 127.0.0.1 --port=13306 < seed/02_insert_data.sql

# for MySQL 8.0 container
mysql --user root --host 127.0.0.1 --port=23306 < seed/01_setup_tables__mysql_8_0.sql
mysql --user root --host 127.0.0.1 --port=23306 < seed/02_insert_data.sql
