#!/bin/bash

# Execute sudo bash ./deletetopo.sh

# ----------------------- DELETE TOPOLOGY STEP --------------------------------

sudo ./sdntool bridge --delete --switchname ovs1 --bridgename br1
sudo ./sdntool bridge --delete --switchname ovs1 --bridgename br2
