#!/bin/bash
echo "Installing Nginx"
sudo apt-get update
sudo apt-get install -y nginx
sudo systemctl enable nginx --now
