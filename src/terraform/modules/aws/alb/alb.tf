resource "aws_alb" "alb" {
  name            = "${substr(var.name, 0, length(var.name) <= 32 ? length(var.name) : 31)}"
  subnets         = ["${var.subnet_ids}"]
  security_groups = ["${var.security_groups}"]
  internal        = "${var.internal}"

  tags {
    Name        = "${var.name}-lb"
    Service     = "${var.name}"
    Environment = "${var.env}"
  }

  access_logs {
    bucket = "${var.log_bucket}"
  }
}

resource "aws_alb_target_group" "target_group" {
  name     = "${substr(var.name, 0, length(var.name) <= 32 ? length(var.name) : 31)}"
  port     = 80
  protocol = "HTTP"
  vpc_id   = "${var.vpc_id}"

  health_check = {
    path = "${var.health_check_endpoint}"

    healthy_threshold   = 2
    unhealthy_threshold = 2
    timeout             = 5
    interval            = 30
    matcher             = "200-299" // Allow any 2xx response pass the healthcheck
  }

  tags {
    Name        = "${var.name}-target-group"
    Service     = "${var.name}"
    Environment = "${var.env}"
  }
}

resource "aws_alb_listener" "external" {
  count = "${! var.internal ? 1 : 0}"

  load_balancer_arn = "${aws_alb.alb.arn}"
  port              = "443"
  protocol          = "HTTPS"
  certificate_arn   = "${var.ssl_certificate_arn}"

  default_action {
    target_group_arn = "${aws_alb_target_group.target_group.arn}"
    type             = "forward"
  }
}

resource "aws_alb_listener" "internal" {
  count = "${var.internal ? 1 : 0}"

  load_balancer_arn = "${aws_alb.alb.arn}"
  port              = "80"
  protocol          = "HTTP"

  default_action {
    target_group_arn = "${aws_alb_target_group.target_group.arn}"
    type             = "forward"
  }
}
