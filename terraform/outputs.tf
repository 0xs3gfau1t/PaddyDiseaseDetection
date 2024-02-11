output "webserver_public_ip" {
  value = azurerm_public_ip.webserver_public_ip.ip_address
}

output "webserver_private_ip" {
  value = azurerm_network_interface.webserver_nic.private_ip_address
}

output "ml_machine_private_ip" {
  value = azurerm_network_interface.ml_machine_nic.private_ip_address
}

output "username" {
  value = var.ssh_username
}
