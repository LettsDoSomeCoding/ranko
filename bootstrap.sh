#!/usr/bin/env bash

apt update
apt install -y docker.io
systemctl enable docker
systemctl start docker
usermod -aG docker vagrant
newgrp docker
docker pull golang