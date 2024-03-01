#!/bin/bash
sudo apt-get update && sudo apt-get upgrade -y
echo "Installing RabbitMQ Server"
sudo apt-get install -y rabbitmq-server
echo "Enable RabbitMQ Server"
sudo systemctl enable rabbitmq-server --now
echo "Enable RabbitMQ Management plugin"
sudo rabbitmq-plugins enable rabbitmq_management # Default port: 15672
echo "Creating new user"
sudo rabbitmqctl add_user siyo 00000
echo "Setting permissions for user in vhost /"
sudo rabbitmqctl set_permissions -p / siyo ".*" ".*" ".*"
sudo systemctl restart rabbitmq-server
