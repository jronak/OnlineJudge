#!/bin/bash
sudo apt-get install lxc debootstrap bridge-utils -y
sudo su
wget http://lxc-webpanel.github.com/tools/install.sh -O - | bash