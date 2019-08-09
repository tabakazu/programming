# Provider
provider "aws" {
  region = "${var.region}"
}

# VPC module
module "network" {
  source = "./modules/network"
  
  namespace         = "${var.namespace}"
  region            = "${var.region}"
  availability_zone = "${var.availability_zone}"
}
