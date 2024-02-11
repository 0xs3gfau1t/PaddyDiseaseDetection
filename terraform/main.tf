terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0.2"
    }
  }
  cloud {
    organization = "0xsegFault"
    workspaces {
      name = "MajorProject"
    }
  }
  required_version = ">= 1.1.0"
}

provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "rg" {
  name     = var.resource_group_name
  location = "centralindia"
  tags = {
    Environment = "MajorProject"
    Team        = "DevOps"
  }
}

resource "azurerm_virtual_network" "vnet" {
  name                = "vnet"
  location            = azurerm_resource_group.rg.location
  address_space       = ["10.0.0.0/16"]
  resource_group_name = azurerm_resource_group.rg.name
}

resource "azurerm_subnet" "subnet" {
  name                 = "subnet"
  resource_group_name  = azurerm_resource_group.rg.name
  virtual_network_name = azurerm_virtual_network.vnet.name
  address_prefixes     = ["10.0.2.0/24"]
}

resource "azurerm_public_ip" "webserver_public_ip" {
  name                = "webserver-public-ip"
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  allocation_method   = "Static"
}

resource "azurerm_network_interface" "webserver_nic" {
  name                = "webserver-nic"
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.subnet.id
    private_ip_address_allocation = "Static"
    private_ip_address            = "10.0.2.100"
    public_ip_address_id          = azurerm_public_ip.webserver_public_ip.id
  }
}

resource "azurerm_linux_virtual_machine" "webserver" {
  name                = "webserver"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  size                = "Standard_B1s"

  admin_username = var.ssh_username

  network_interface_ids = [azurerm_network_interface.webserver_nic.id]

  custom_data = base64encode(file("scripts/webserver.sh"))

  admin_ssh_key {
    username   = var.ssh_username
    public_key = var.ssh_public_key
  }
  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }
  source_image_reference {
    version   = "latest"
    sku       = "11-backports-gen2"
    offer     = "debian-11"
    publisher = "Debian"
  }
}

resource "azurerm_postgresql_server" "backend_postgres_server" {
  name                             = "segfault-postgresql-server"
  location                         = azurerm_resource_group.rg.location
  resource_group_name              = azurerm_resource_group.rg.name
  sku_name                         = "B_Gen5_2"
  storage_mb                       = 5120
  backup_retention_days            = 7
  geo_redundant_backup_enabled     = false
  auto_grow_enabled                = false
  administrator_login              = var.ssh_username
  administrator_login_password     = var.postgres_password
  ssl_enforcement_enabled          = false
  ssl_minimal_tls_version_enforced = "TLSEnforcementDisabled"
  version                          = 9.5

}

resource "azurerm_postgresql_database" "backend_db" {
  name                = "postgresql"
  resource_group_name = azurerm_resource_group.rg.name
  server_name         = azurerm_postgresql_server.backend_postgres_server.name
  charset             = "UTF8"
  collation           = "English_United States.1252"
}

resource "azurerm_network_interface" "ml_machine_nic" {
  name                = "ml-machine-nic"
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.subnet.id
    private_ip_address_allocation = "Static"
    private_ip_address            = "10.0.2.200"
  }
}

resource "azurerm_linux_virtual_machine" "ml_server" {
  name                = "ml-server"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  size                = "Standard_B1s"

  admin_username = var.ssh_username

  network_interface_ids = [azurerm_network_interface.ml_machine_nic.id]

  admin_ssh_key {
    username   = var.ssh_username
    public_key = var.ssh_public_key
  }
  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }
  source_image_reference {
    version   = "latest"
    sku       = "11-backports-gen2"
    offer     = "debian-11"
    publisher = "Debian"
  }
}

resource "azurerm_network_security_group" "allowed_ports" {
  name                = "allowed_ports"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location

  security_rule {
    name                       = "http"
    priority                   = 100
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "3000"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "https"
    priority                   = 200
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "443"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "ssh"
    priority                   = 300
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "22"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }
}
