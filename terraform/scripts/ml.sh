#!/bin/bash

echo "Preparing dependencies"
sudo apt-get update && sudo apt-get upgrade -y
sudo apt-get install -y python3 python3-pip git

cd /home/$USER/paddydiseasedetection/ai
sudo pip install -r requirements.txt --no-cache-dir
nohup bash -c "while :; do python3 main.py; done &" # restart this process everytime it's killed due to various reasons
