module "vpc" {
  source = "./vpc"

  env = "${var.env}"
}

module "subnets" {
  source = "./subnets"

  env                = "${var.env}"
  vpc_id             = "${module.vpc.id}"
  vpc_cidr           = "${module.vpc.cidr}"
  availability_zones = "${var.availability_zones}"
}

module "bastion" {
  source = "./bastion"

  region             = "${var.region}"
  env                = "${var.env}"
  vpc_id             = "${module.vpc.id}"
  availability_zones = "${var.availability_zones}"
  public_subnet_ids  = "${module.subnets.public_subnet_ids}"
  key_name           = "${var.key_name}"
}

output "vpc_id" {
  value = "${module.vpc.id}"
}

output "public_subnet_ids" {
  value = ["${module.subnets.public_subnet_ids}"]
}

output "private_subnet_ids" {
  value = ["${module.subnets.private_subnet_ids}"]
}

output "bastion_security_group_id" {
  value = "${module.bastion.security_group_id}"
}