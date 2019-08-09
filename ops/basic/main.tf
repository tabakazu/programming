# Provider
provider "aws" {
  region = "${var.region}"
}

# Network module
module "network" {
  source = "./modules/network"
  
  namespace         = "${var.namespace}"
  region            = "${var.region}"
  availability_zone = "${var.availability_zone}"
}

# Web module
module "web" {
  source = "./modules/web"
  
  namespace          = "${var.namespace}"
  ami                = "${var.ami}"
  public_key         = "${var.public_key}"
  vpc_id             = "${module.network.vpc_id}"
  subnet_public_a_id = "${module.network.subnet_public_a_id}"
  subnet_public_c_id = "${module.network.subnet_public_c_id}"
}
