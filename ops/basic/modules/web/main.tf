# EC2
resource "aws_instance" "web_a" {
  ami                         = "${var.ami}"
  subnet_id                   = "${var.subnet_public_a_id}"
  instance_type               = "t2.micro"
  private_ip                  = "10.0.0.100"
  associate_public_ip_address = "true"
  key_name                    = "${aws_key_pair.this.key_name}"
  
  vpc_security_group_ids      = [
    "${aws_security_group.this.id}"
  ]

  tags = {
    Name = "${var.namespace}-web-01"
  }
}

# Key pair
resource "aws_key_pair" "this" {
  key_name   = "${var.namespace}-web"
  public_key = "${var.public_key}"
}

# Security Group
resource "aws_security_group" "this" {
  name        = "${var.namespace}-web"
  vpc_id      = "${var.vpc_id}"

  ingress {
    protocol    = "tcp"
    from_port   = 22
    to_port     = 22
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    protocol    = "-1"
    from_port   = 0
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }
}