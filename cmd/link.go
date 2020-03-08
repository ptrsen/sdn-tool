/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/ptrsen/sdn-tool/internal/pkg/system"
	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Create link between bridge and host",
	Long: `Create link between bridge and host. For example:
			sudo ./sdntool link --create --switchname ovs1 --bridge br1 --host h1
		    sudo ./sdntool link --delete --switchname ovs1 --bridge br1 --host h1`,
	Run: link,
}

var linkCreate bool
var linkDelete bool

var hostname string
var iphost string

func init() {
	rootCmd.AddCommand(linkCmd)
	linkCmd.PersistentFlags().BoolVar(&linkCreate, "create",false,"Create link between bridge and host")
	linkCmd.PersistentFlags().BoolVar(&linkDelete, "delete",false,"Delete link between bridge and host")

	linkCmd.Flags().StringVarP(&swname, "switchname", "s", "", "Container switch name")
	linkCmd.Flags().StringVarP(&brname, "bridge", "b", "", "Bridge name ")
	linkCmd.Flags().StringVarP(&hostname, "host", "c", "", "Host name")
	linkCmd.Flags().StringVarP(&iphost, "iphost", "i", "", "ip for Host")

}


func link (cmd *cobra.Command, args []string) {
   // Generate Random MAC
  //  fmt.Println( system.GenerateMac().String() )

 // GET PID
 /*
	PullCmds := []string{
		"sudo docker inspect -f '{{.State.Pid}}' ovs1",
	}
	for _, cmd := range PullCmds {
		ovsPid, err := system.ShellExecOutput("/bin", "sh", "-c", cmd)
		system.CheckError("Cannot get switch pid", err)
		fmt.Println(ovsPid)
	}
*/


	// Create link
	if linkCreate{

		// Getting switch container process id
		cmd := "sudo docker inspect -f '{{.State.Pid}}' "+ swname
		ovsPid, err := system.ShellExecOutput("/bin", "sh", "-c", cmd)
		system.CheckError("Cannot get switch pid", err)

		// Getting host container process id
		cmd = "sudo docker inspect -f '{{.State.Pid}}' "+ hostname
		hostPid, err := system.ShellExecOutput("/bin", "sh", "-c", cmd)
		system.CheckError("Cannot get switch pid", err)

		// Create netns directory , and Soft link (symlink) process network namespace from /proc directory into the /var/run directory
		PullCmds := []string{
			"sudo mkdir -p /var/run/netns",
			"sudo ln -sf /proc/" + ovsPid + "/ns/net \"/var/run/netns/" + ovsPid + "\" ",
			"sudo ln -sf /proc/" + hostPid + "/ns/net \"/var/run/netns/" + hostPid + "\" ",
		}
		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot create soft link for containers",err)
		}

		// Create a veth virtual-interface pair and Assign the interfaces to corresponding namespaces
		PullCmds = []string{
			"sudo ip link add " + hostname + "-tap" + " type veth peer name " + brname + "-port-" + hostname,
			"sudo ip link set "+ brname + "-port-" + hostname + " netns "+ ovsPid,
			"sudo ip link set " + hostname + "-tap" + " netns "+ hostPid,
		}
		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot create veth and assign to container namespaces",err)
		}

		// OVS switch
		// Add port to ovs bridge and attach interface, up interface
		PullCmds = []string{
			"sudo docker exec " + swname +" ovs-vsctl --may-exist add-port " + brname + " " + brname + "-port-" + hostname,
			"sudo ip netns exec " + ovsPid +" ip link set dev " + brname + "-port-" + hostname + " up",
		}
		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot add port to container switch and setting up",err)
		}

		// HOST
		// Change interface name , Assing Mac address , Setting interface up , Assing ip address
		PullCmds = []string{
			"sudo ip netns exec " + hostPid +" ip link set dev " + hostname + "-tap" +" name eth0",
			"sudo ip netns exec " + hostPid +" ip link set eth0 address " + system.GenerateMac().String() ,
			"sudo ip netns exec " + hostPid +" ip link set eth0 up" ,
			"sudo ip netns exec " + hostPid +" ip addr add " + iphost +" dev eth0 " ,
		}
		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot setting network container host up",err)
		}


	}




	// Delete link
	if linkDelete {

		// Getting switch container process id
		cmd := "sudo docker inspect -f '{{.State.Pid}}' "+ swname
		ovsPid, err := system.ShellExecOutput("/bin", "sh", "-c", cmd)
		system.CheckError("Cannot get switch pid", err)

		// Getting host container process id
		cmd = "sudo docker inspect -f '{{.State.Pid}}' "+ hostname
		hostPid, err := system.ShellExecOutput("/bin", "sh", "-c", cmd)
		system.CheckError("Cannot get switch pid", err)

		PullCmds := []string{
			"sudo docker exec " + swname + " " + "ovs-vsctl --if-exist del-port " + brname + "-port-" + hostname,
			"sudo rm -r /var/run/netns/" + hostPid,
			"sudo rm -r /var/run/netns/" + ovsPid,
		}

		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot delete link between bridge and host ",err)
		}

	}

}



