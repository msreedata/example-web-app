#!/bin/bash
:' Run below lines manually first -- this will build the docker image from dockerFile and start the container

cd $HOME
mkdir -p dockerstore
cd dockerstore/
git clone https://github.com/msreedata/hello-Go.git
cd $HOME/dockerstore/hello-Go

'

cd $HOME/dockerstore/hello-Go
sudo yum -y install docker
ls
read -rsp $'Press any key to continue...\n' -n1 key

systemctl status docker
systemctl start docker
systemctl status docker
systemctl enable docker
systemctl status docker
read -rsp $'Press any key to continue...\n' -n1 key


sudo docker build . -t hellogo:v2
sudo docker image ls
read -rsp $'Press any key to continue...\n' -n1 key

sudo docker run --rm -d -p 82:80 --name hgo2 hellogo:v2
sudo docker ps -a
#
