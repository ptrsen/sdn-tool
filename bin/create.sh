#!/bin/bash

# Execute sudo bash ./create.sh


# ------------------ INSTALLATION ---------------------------------

# Installs docker (just the first time)
# sudo ./sdntool install --docker

# -----------------------------------------------------------------


# ------------------ CREATE STEP ---------------------------------

# Create docker images from Dockerfiles
sudo ./sdntool images --create

# Create SDN control Network  (communication between controller and Switches)
sudo ./sdntool create --network --name control --subnet 172.10.0.0/16 --iprange 172.10.1.0/24 --gateway 172.10.1.1

# CREATE NETWORK COMPONENTS
#Create Controller (ONOS)
sudo ./sdntool create --controller --name ctl1 --netname control --ip 172.10.1.2
# Create OVS Switch (just need one for several bridges, its possible to create more )
sudo ./sdntool create --switch --name ovs1 --netname control --ip 172.10.1.3
# Create 2 Host
sudo ./sdntool create --host --name h1
sudo ./sdntool create --host --name h2

# -----------------------------------------------------------------