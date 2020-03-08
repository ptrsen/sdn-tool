#!/bin/bash

# Execute sudo bash ./delete.sh

# --------------------- DELETE STEP --------------------------------


# DELETE COMPONENTS
# Delete Host
sudo ./sdntool delete --host --name h1
sudo ./sdntool delete --host --name h2

# Delete Ovs Switch
sudo ./sdntool delete --switch --name ovs1

#Delete ONOS Controller
sudo ./sdntool delete --controller --name ctl1

# Delete SDN control Network  (communication between controller and Switches)
sudo ./sdntool delete --network --name control

# Deletes images ( to destroy everthing )
# sudo ./sdntool images --delete

# -----------------------------------------------------------------