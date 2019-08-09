# RDS
resource "aws_db_instance" "this" {
  identifier           = "${var.namespace}-db"
  skip_final_snapshot  = true
  allocated_storage    = 10
  storage_type         = "gp2"
  engine               = "mysql"
  engine_version       = "5.7"
  instance_class       = "db.t2.micro"
  name                 = "mydb"
  username             = "dbuser"
  password             = "passw0rd"
  db_subnet_group_name = "${aws_db_subnet_group.this.name}"
  parameter_group_name = "${aws_db_parameter_group.default.name}"
  security_group_names = [
    "${aws_security_group.this.name}"
  ]
}

# DB Subnet Group
resource "aws_db_subnet_group" "this" {
  name       = "${var.namespace}-db"
  subnet_ids = [
    "${var.subnet_private_a_id}",
    "${var.subnet_private_c_id}"
  ]
}

resource "aws_db_parameter_group" "default" {
  name   = "${var.namespace}-db"
  family = "mysql5.7"

  parameter {
    name  = "character_set_server"
    value = "utf8"
  }

  parameter {
    name  = "character_set_client"
    value = "utf8"
  }
}

# Security Group
resource "aws_security_group" "this" {
  name   = "${var.namespace}-db"
  vpc_id = "${var.vpc_id}"

  ingress {
    protocol        = "tcp"
    from_port       = 3306
    to_port         = 3306
    security_groups = ["${var.security_group_ec2_id}"]
  }

  egress {
    protocol    = "-1"
    from_port   = 0
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }
}