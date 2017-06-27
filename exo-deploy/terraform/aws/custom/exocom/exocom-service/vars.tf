/* Variables */

variable "elb_subnet_ids" {
  type        = "list"
  description = "List of public or private ID's the ALB should live in"
}

variable "cluster_id" {
  description = "ID of the ECS cluster"
}

variable "command" {
  description = "Starting command to run in container"
  type = "list"
}

variable "container_port" {
  description = "Port number on the container to bind the ALB to"
  default     = 80
}

variable "cpu_units" {
  description = "Number of cpu units to reserve for the container"
}

variable "docker_image" {
  description = "ECS repository URI of Docker image"
}

variable "ecs_role_arn" {
  description = "ARN of the ECS IAM role"
}

variable "env" {
  description = "Name of the environment, used for naming and prefixing"
}

variable "environment_variables" {
  type        = "map"
  description = "Environment variables to pass to a container"
}

variable "health_check_endpoint" {
  description = "Endpoint for the elb to hit when performing health checks"
  default     = "/"
}

variable "memory_reservation" {
  description = "Soft limit (in MiB) of memory to reserve for the container"
}

variable "name" {
  description = "Name of the service"
}

variable "region" {
  description = "Region of the environment, for example, us-west-2"
}

variable "security_groups" {
  description = "IDs of security groups (should include exocom external ALB and exocom cluster)"
  type        = "list"
}

variable "vpc_id" {
  description = "ID of the VPC"
}

/* Output */

output "url" {
  value = "${module.elb.url}"
}