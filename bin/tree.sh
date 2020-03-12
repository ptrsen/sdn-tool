#!/bin/bash
# Execute sudo bash ./tree.sh




# ------------------ CREATE STEP ---------------------------------

# Create docker images from Dockerfiles
#sudo ./sdntool images --create

# Create SDN control Network  (communication between controller and Switches)
#sudo ./sdntool create --network --name control --subnet 172.10.0.0/16 --iprange 172.10.1.0/24 --gateway 172.10.1.1

# CREATE NETWORK COMPONENTS
#Create Controller (ONOS)
#sudo ./sdntool create --controller --name c1 --netname control --ip 172.10.1.2
# Create OVS Switch (just need one for several bridges, its possible to create more )
#sudo ./sdntool create --switch --name ovs --netname control --ip 172.10.1.3


# Create 2 Host
#sudo ./sdntool create --host --name h1
#sudo ./sdntool create --host --name h2

# -----------------------------------------------------------------



# ------------------ TOPO STEP ---------------------------------
totalBridges=3

for (( i=1; i<=$totalBridges; i++ ))
  do
    echo "Creating ovs-br$i"
    #sudo ./sdntool bridge --create --switchname ovs --bridgename ovs-br$i --controllerip tcp:172.10.1.2:6633
 done



for (( i=1; i<=$totalBridges; i++ ))
  do
    echo "--- ovs-br$i"
    echo "--- ovs-br$(($i+1))"
    #sudo ./sdntool patch --create --switchname ovs --bridgeA ovs-br$i --bridgeB ovs-br$i
  done



# -----------------------------------------------------------------


