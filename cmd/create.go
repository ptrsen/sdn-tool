/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions andN
limitations under the License.
*/
package cmd

import (
	"github.com/ptrsen/sdn-tool/internal/pkg/system"
	"os"

	"github.com/spf13/cobra"
)


// Controllers
var ctlogsLocalDirectory = "containers/controllers/log"  // Logs
var ctllogsContainerDirectory = "/app/log"
var ctlconfLocalDirectory = "containers/controllers/conf" // configuration
var ctlconfContainerDirectory = "/app/conf"


// Switch
var swlogsLocalDirectory = "containers/switches/log"  // Logs
var swlogsContainerDirectory = "/app/log"
var swconfLocalDirectory = "containers/switches/conf" // configuration
var swconfContainerDirectory = "/app/conf"

// Host
var hostlogsLocalDirectory = "containers/hosts/log"  // Logs
var hostlogsContainerDirectory = "/app/log"
var hostconfLocalDirectory = "containers/hosts/conf" // configuration
var hostconfContainerDirectory = "/app/conf"


// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create container based network devices",
	Long: `Create container based network devices. For example:
			 sudo ./sdntool create --network --name control --subnet 172.10.0.0/16 --iprange 172.10.1.0/24 --gateway 172.10.1.1
             sudo ./sdntool create --controller --name ctl1 --netname control --ip 172.10.1.2 
             sudo ./sdntool create --switch --name ovs1 --netname control --ip 172.10.1.3
			 sudo ./sdntool create --host --name h1 `,
	Run: create,
}

// Flags
var ctlCreate bool  // --Controller flag
var swCreate bool   // --Switch flag
var hCreate bool    //  --Host flag
var netCreate bool  // --Network flag


var devName string  // --name flag
var subNet string   // --subnet flag
var ipRange string  // --iprange flag
var gateWay string  // --gateway flag

var netName string  // --network flag
var ip4net string  // --ip flag

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().BoolVar(&netCreate, "network",false,"Create docker network")
	createCmd.PersistentFlags().BoolVar(&ctlCreate, "controller",false,"Create docker controller device")
	createCmd.PersistentFlags().BoolVar(&swCreate, "switch",false,"Create docker switch device")
	createCmd.PersistentFlags().BoolVar(&hCreate, "host",false,"Create docker host device")


	createCmd.Flags().StringVarP(&devName, "name", "n", "", "Device name")
	createCmd.Flags().StringVarP(&subNet, "subnet", "s", "", "Subnet")
	createCmd.Flags().StringVarP(&ipRange, "iprange", "i", "", "ip range")
	createCmd.Flags().StringVarP(&gateWay, "gateway", "g", "", "Gateway")


	createCmd.Flags().StringVarP(&netName, "netname", "t", "", "network name")
	createCmd.Flags().StringVarP(&ip4net, "ip", "p", "", "Ipv4")

}

func create(cmd *cobra.Command, args []string) {

	projectPath, er := os.Getwd()
	system.CheckError("Cant find Project path", er)


	// Create Network
	if netCreate {

		// Create  network
		PullCmds := []string{
			"sudo docker network create --driver=bridge" + " " +
				"--subnet="+ subNet + " " +
				"--ip-range="+ ipRange + " " +
				"--gateway=" + gateWay + " " +
				devName,
		}

		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot create Network ",err)
		}

	}


	// Create Host
	if hCreate {

		// Create directory for log files , if not exists
		if _, err := os.Stat(hostlogsLocalDirectory + "/" + devName); os.IsNotExist(err) {
			err := os.MkdirAll(hostlogsLocalDirectory +"/"+ devName, 0777)
			system.CheckError("Unable to create containers/switches/log/deviceName directory", err)
		}
		// Create directory for configuration files , if not exists
		if _, err := os.Stat( hostconfLocalDirectory + "/" + devName); os.IsNotExist(err) {
			err := os.MkdirAll(hostconfLocalDirectory +"/"+ devName, 0777)
			system.CheckError("Unable to create containers/switches/conf/deviceName directory", err)
		}

		// Create host
		PullCmds := []string{
			"sudo docker run -itd" + " " +
				"--network none " + " " +
				"--name " + devName + " " +
				"--volume "+ projectPath + "/" + hostlogsLocalDirectory + ":"+ hostlogsContainerDirectory + " " +
				"--volume "+ projectPath + "/" + hostconfLocalDirectory + ":"+ hostconfContainerDirectory + " " +
				"host",
		}

		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot create host container",err)
		}
	}

	// Create Switch
	if swCreate {

		// Create directory for log files , if not exists
		if _, err := os.Stat(swlogsLocalDirectory + "/" + devName); os.IsNotExist(err) {
			err := os.MkdirAll(swlogsLocalDirectory +"/"+ devName, 0777)
			system.CheckError("Unable to create containers/switches/log/deviceName directory", err)
		}
		// Create directory for configuration files , if not exists
		if _, err := os.Stat( swconfLocalDirectory + "/" + devName); os.IsNotExist(err) {
			err := os.MkdirAll(swconfLocalDirectory +"/"+ devName, 0777)
			system.CheckError("Unable to create containers/switches/conf/deviceName directory", err)
		}

		// Create OVS switch
		PullCmds := []string{
			"sudo docker run -itd"+ " " +
				"--network "+ netName + " " +
				"--ip "+ ip4net + " " +
				"--cap-add NET_ADMIN" + " " +
				"--name " + devName + " " +
				"--volume "+ projectPath + "/" + swlogsLocalDirectory + ":"+ swlogsContainerDirectory + " " +
				"--volume "+ projectPath + "/" + swconfLocalDirectory + ":"+ swconfContainerDirectory + " " +
				"switch",
		}

		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot create switch container ",err)
		}

	}

	// Create controller
	if ctlCreate {

		// Create directory for log files , if not exists
		if _, err := os.Stat(ctlogsLocalDirectory + "/" + devName); os.IsNotExist(err) {
			err := os.MkdirAll(ctlogsLocalDirectory+"/"+ devName, 0777)
			system.CheckError("Unable to create containers/controllers/log/deviceName directory", err)
		}
		// Create directory for configuration files , if not exists
		if _, err := os.Stat(ctlconfLocalDirectory + "/" + devName); os.IsNotExist(err) {
			err := os.MkdirAll(ctlconfLocalDirectory+"/"+ devName, 0777)
			system.CheckError("Unable to create containers/controllers/conf/deviceName directory", err)
		}

		// Create Onos controller
		PullCmds := []string{
			"sudo docker run -itd"+ " " +
				"--name " + devName + " " +
				"--network "+ netName + " " +
				"--ip "+ ip4net + " " +
				"--volume "+ projectPath + "/" + ctlogsLocalDirectory + ":"+ ctllogsContainerDirectory + " " +
				"--volume "+ projectPath + "/" + ctlconfLocalDirectory + ":"+ ctlconfContainerDirectory + " " +
				"controller",
		}

		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot create controller container",err)
		}

	}



}
