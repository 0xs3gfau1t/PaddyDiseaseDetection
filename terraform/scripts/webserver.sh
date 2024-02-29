#!/bin/bash
echo "Installing Nginx"
sudo apt-get update && sudo apt-get upgrade
sudo apt-get install -y nginx
sudo systemctl enable nginx --now

echo "Fetch go binary"
sudo apt-get install -y golang-go
wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
echo "Extracting go archive"
sudo tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
echo "Setting GOPATH"
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc

echo "Generating db functions"
cd ~/paddydiseasedetection/backend
go run -mod=mod entgo.io/ent/cmd/ent generate .ent/schema
echo "Installing api binary"
go build
echo "Installation complete. Setup Nginx now."

read -p "Do you want to setup rabbitmq port forwarding? (y/n) => " optn
if [[ "$optn" != "y" ]];then
  exit
fi

# Setup port forward with socat
echo "Setting up port forwarding for rabbitmq management"
sudo apt-get install socat -y
sudo socat tcp-listen:3000,fork tcp:10.0.2.201:15672&
disown
