#!/bin/bash

# This script install latest golang version 
# https://golang.org/doc/install
# Execute: bash go-install.sh 

# Setting git 
gitUser=ptrsen
gitRepo=sdn-tool.git
gitBranch=master 

# Getting Golang latest version and url 
url="$(wget -qO- https://golang.org/dl/ | grep -oP 'https:\/\/dl\.google\.com\/go\/go([0-9\.]+)\.linux-amd64\.tar\.gz' | head -n 1 )"
latest="$(echo "$url" | grep -oP 'go[0-9\.]+' | grep -oP '[0-9\.]+' | head -c -2 )"

# Update enviroment
cd ~  || exit
echo -e "\n\n Updating enviroment... \n" 
sudo apt update
sudo apt -y dist-upgrade

# install pre-reqs
echo -e "\n\n Installing Pre-requirements ... \n"
sudo apt -y install python3 git curl wget

# Remove old Golang version 
echo -e "\n\n Removing old golang ... \n"
sudo rm -rf /usr/local/go
sudo rm -rf "$HOME"/go

# install latest Golang version
echo -e "\n\n Installing golang $latest ... \n"
# Download and extract golang , default home directory  path /usr/local/go 
wget "$url"
sudo tar -xvf go"$latest".linux-amd64.tar.gz
rm -rf go"$latest".linux-amd64.tar.gz
sudo chown -R root:root ./go
sudo mv go /usr/local

# Setting GO Paths
{
echo "export GOROOT=/usr/local/go"    #use Default, if the installion is done in $HOME/go  GOROOT=$HOME/go"
echo "export GOPATH=$HOME/go"
echo "export PATH=$HOME/go/bin:/usr/local/go/bin:$PATH"
}>> ~/.bashrc

. ~/.bashrc

# Setting Directory 
mkdir -p "$HOME"/go/src/github.com/$gitUser  # end
echo -e "\n\n Work directory $HOME/go/src/github.com/$gitUser ... \n"
cd "$HOME"/go/src/github.com/$gitUser || exit
git clone -b $gitBranch https://github.com/$gitUser/$gitRepo
cd ~  || exit
. ~/.bashrc
cd $HOME/go/src/github.com/$gitUser # project workspace

# Checking Version
go version

