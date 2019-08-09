output "vpc_id" {
  value = "${aws_vpc.this.id}"
}

output "subnet_public_a_id" {
  value = "${aws_subnet.public_a.id}"
}

output "subnet_public_c_id" {
  value = "${aws_subnet.public_c.id}"
}

output "subnet_private_a_id" {
  value = "${aws_subnet.private_a.id}"
}

output "subnet_private_c_id" {
  value = "${aws_subnet.private_c.id}"
}