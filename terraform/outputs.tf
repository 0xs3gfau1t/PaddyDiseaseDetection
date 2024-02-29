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

output "rabbitmq_private_ip" {
  value = azurerm_linux_virtual_machine.rabbitmq_server.private_ip_address
}

output "dababase_private_dns" {
  value = azurerm_private_dns_zone.db_private_dns_zone.name
}

output "database_connection_uri" {
 value = azurerm_postgresql_flexible_server.db_server.name 
}
