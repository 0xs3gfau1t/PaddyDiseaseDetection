#!/bin/bash
sudo apt-get update
sudo apt-get install -y rabbitmq-server
sudo systemctl enable rabbitmq-server --now
sudo rabbitmq-plugins enable rabbitmq_management --now
