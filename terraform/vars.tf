variable "resource_group_name" {
  default = "MajorProjectRG"
}

variable "ssh_username" {
  default = "siyo"
}

variable "postgres_password" {
  default = "Postgres00000*"
}

variable "ssh_public_key" {
  default = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDZX9G2Q4NJTFChxY55CL5ItKP3VcOSO4B9/TD5rtR+3kYDa6X5L7EEH76Yo87EFB32EDA81mKtBs4jeM7Bb+UAR4BU6kOdVRTNfljsNQv+DZjxn2vaqJoW+F+ik5rhKYgTTaSRmAZWaAJq05YRvgKkVRS5YbYp+O/FnioUPnoU2xnUCLZXc/ps8xCXrL4/yG7RtZ1EH84H8XS5PiLlLuOWQ3OZGWtXUtA+Jwmlsz5HIraeEY+OViXMgpLrRKqxOWaQVZlD6/4xY+utUpsi/QGRKPOPTaC+0uDs7g14FikvArpoF+2gnGfg9wIPDADJPs415W37CnfJKD9T6oeFqDyUCIdBYWHqeUMT4kY/28NObIVF9KWkCIwYNIL4skEJ5GYfZOH3jxUzZPjJqvwlE45c+5Zr/UhOpjuc9bUg360E4Bi0MYZlMOkisBfS0+VBin7w55XMDoV31IjJcBxAjeGU5hBrD+V20NwQU7irPWMLPdYwU5WKrQZ8cZDY/YRQHCc= siyo@arch-sama"
}

variable "mq_private_ip" {
  default = "10.0.2.201"
}

variable "webserver_private_ip" {
  default = "10.0.2.100"
}

variable "webserver_db_private_ip" {
  default = "10.0.3.100"
}
variable "ml_server_private_ip" {
  default = "10.0.2.200"
}

variable "mq_admin_password" {
  default = "00000"
}

variable "postgresdb_name" {
  default = "postgresdb-psqlflexibleserver"
}
