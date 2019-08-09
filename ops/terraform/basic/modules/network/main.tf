# VPC
resource "aws_vpc" "this" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = "true"
  enable_dns_hostnames = "true"

  tags = {
    Name = "${var.namespace}"
  }
}

# Internet Gateway
resource "aws_internet_gateway" "this" {
  vpc_id = "${aws_vpc.this.id}"

  tags = {
    Name = "${var.namespace}"
  }
}

# Public Subnet
resource "aws_subnet" "public_a" {
  vpc_id            = "${aws_vpc.this.id}"
  availability_zone = "${var.availability_zone["a"]}"
  cidr_block        = "10.0.0.0/24"

  tags = {
    Name = "${var.namespace}-public-a"
  }
}

resource "aws_subnet" "public_c" {
  vpc_id            = "${aws_vpc.this.id}"
  availability_zone = "${var.availability_zone["c"]}"
  cidr_block        = "10.0.1.0/24"
  tags = {
    Name = "${var.namespace}-public-c"
  }
}

# Public Route Table
resource "aws_route_table" "public" {
  vpc_id = "${aws_vpc.this.id}"

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = "${aws_internet_gateway.this.id}"
  }

  tags = {
    Name = "${var.namespace}-public"
  }
}

# Public Subnet Route Table Association
resource "aws_route_table_association" "public_a" {
  subnet_id      = "${aws_subnet.public_a.id}"
  route_table_id = "${aws_route_table.public.id}"
}

resource "aws_route_table_association" "public_c" {
  subnet_id      = "${aws_subnet.public_c.id}"
  route_table_id = "${aws_route_table.public.id}"
}

# Private Subnet
resource "aws_subnet" "private_a" {
  vpc_id            = "${aws_vpc.this.id}"
  availability_zone = "${var.availability_zone["a"]}"
  cidr_block        = "10.0.2.0/24"

  tags = {
    Name = "${var.namespace}-private-a"
  }
}

resource "aws_subnet" "private_c" {
  vpc_id            = "${aws_vpc.this.id}"
  availability_zone = "${var.availability_zone["c"]}"
  cidr_block        = "10.0.3.0/24"
  tags = {
    Name = "${var.namespace}-private-c"
  }
}

# Private Route Table
resource "aws_route_table" "private" {
  vpc_id = "${aws_vpc.this.id}"

  tags = {
    Name = "${var.namespace}-private"
  }
}

# Private Subnet Route Table Association
resource "aws_route_table_association" "private_a" {
  subnet_id      = "${aws_subnet.private_a.id}"
  route_table_id = "${aws_route_table.private.id}"
}

resource "aws_route_table_association" "private_c" {
  subnet_id      = "${aws_subnet.private_c.id}"
  route_table_id = "${aws_route_table.private.id}"
}