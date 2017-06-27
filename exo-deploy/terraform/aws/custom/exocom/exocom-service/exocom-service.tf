module "elb" {
  source                = "./elb"

  env                   = "${var.env}"
  health_check_endpoint = "${var.health_check_endpoint}"
  name                  = "${var.name}"
  security_groups       = ["${var.security_groups}"]
  subnet_ids            = "${var.elb_subnet_ids}"
  vpc_id                = "${var.vpc_id}"
}

module "task_definition" {
  source                = "./ecs-task-definition"

  command               = "${var.command}"
  container_port        = "${var.container_port}"
  cpu_units             = "${var.cpu_units}"
  docker_image          = "${var.docker_image}"
  env                   = "${var.env}"
  environment_variables = "${var.environment_variables}"
  memory_reservation    = "${var.memory_reservation}"
  name                  = "${var.name}"
  region                = "${var.region}"
}

resource "aws_ecs_service" "service" {
  name                               = "${var.name}"

  cluster                            = "${var.cluster_id}"
  depends_on                         = ["module.elb"]
  deployment_minimum_healthy_percent = 100
  desired_count                      = 1
  iam_role                           = "${var.ecs_role_arn}"
  task_definition                    = "${module.task_definition.task_arn}"

  load_balancer {
    elb_name         = "${module.elb.name}"
    container_name   = "${var.name}"
    container_port   = "${var.container_port}"
  }
}