#!/bin/bash

# Execute sudo bash ./createtopo.sh

# ----------------------- CREATE TOPOLOGY STEP --------------------------------


sudo ./sdntool bridge --create --switchname ovs1 --bridgename br1 --controllerip tcp:172.10.1.2:6633
sudo ./sdntool bridge --create --switchname ovs1 --bridgename br2 --controllerip tcp:172.10.1.2:6633


sudo ./sdntool patch --create --switchname ovs1 --bridgeA br1 --bridgeB br2
#sudo ./sdntool patch --delete --switchname ovs1 --bridgeA br1 --bridgeB br2


sudo ./sdntool link --create --switchname ovs1 --bridge br1 --host h1 --iphost 10.0.0.1/16
sudo ./sdntool link --create --switchname ovs1 --bridge br2 --host h2 --iphost 10.0.0.2/16

#sudo ./sdntool link --delete --switchname ovs1 --bridge br1 --host h1  //revisar
#sudo ./sdntool link --delete --switchname ovs1 --bridge br2 --host h2


