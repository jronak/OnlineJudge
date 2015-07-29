#!/bin/bash
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit 1
fi
cd go/src
cp lxcmain /home/ubuntu
cp start* /home/ubuntu
echo "Run the startJudge in $HOME as user"