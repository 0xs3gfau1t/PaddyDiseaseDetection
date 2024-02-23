variable "resource_group_name" {
  default = "MajorProjectRG"
}

variable "ssh_username" {
}

variable "postgres_password" {
  sensitive = true
}

variable "ssh_public_key" {
  sensitive = true
}

variable "mq_private_ip" {
  default = "10.0.2.201"
}

variable "webserver_private_ip" {
  default = "10.0.2.100"
}

variable "ml_server_private_ip" {
  default = "10.0.2.200"
}

variable "mq_admin_password" {
  sensitive = true
}
