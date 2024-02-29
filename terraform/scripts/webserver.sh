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
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
sudo apt-get remove -y golang-go
sudo apt-get -y autoremove
source ~/.bashrc

echo "Generating db functions"
cd ~/paddydiseasedetection/backend
go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
echo "Installing api binary"
go build
echo "Spawning 5 backend processed"
for i in {8000..8005};do
  ./PaddyDiseaseDetection server -p $i &
done
echo "Installation complete. Setting up Nginx now."
echo """
user www-data;
worker_processes auto;
pid /run/nginx.pid;

events {}

http {
  upstream backend {
    server 127.0.0.1:8000;
    server 127.0.0.1:8001;
    server 127.0.0.1:8002;
    server 127.0.0.1:8003;
  }

  server {
    listen 80;
    location / {
      proxy_pass http://backend;
    }
  }
}
"""


# Setup port forward with socat
echo "Setting up port forwarding for rabbitmq management"
sudo apt-get install socat -y && sudo socat tcp-listen:3000,fork tcp:10.0.2.201:15672&
disown
