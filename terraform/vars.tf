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
