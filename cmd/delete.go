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
	"os"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete container based network devices",
	Long: `Delete container based network devices. For example:
			 sudo ./sdntool delete --network --name control 
             sudo ./sdntool delete --controller --name ctl1  
             sudo ./sdntool delete --switch --name ovs1 
			 sudo ./sdntool delete --host --name h1`,
	Run: delete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.PersistentFlags().BoolVar(&netCreate, "network",false,"Delete docker network")
	deleteCmd.PersistentFlags().BoolVar(&ctlCreate, "controller",false,"Delete docker controller device")
	deleteCmd.PersistentFlags().BoolVar(&swCreate, "switch",false,"Delete docker switch device")
	deleteCmd.PersistentFlags().BoolVar(&hCreate, "host",false,"Delete docker host device")

	deleteCmd.Flags().StringVarP(&devName, "name", "n", "", "Device name")

}

func delete (cmd *cobra.Command, args []string) {

	projectPath, er := os.Getwd()
	system.CheckError("Cant find Project path", er)

	// Delete Network
	if netCreate {

		PullCmds := []string{
			"sudo docker network rm" + " " +
				devName,
		}

		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot delete Network",err)
		}

	}


	// Delete Host
	if hCreate {

		// Delete host
		PullCmds := []string{
			"sudo rm -rf "+ projectPath + "/" + hostlogsLocalDirectory,
			"sudo rm -rf "+ projectPath + "/" + hostconfLocalDirectory,
			"sudo docker container stop" + " " +
				devName,
			"sudo docker container rm" + " " +
				devName,
		}

		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot stop and delete host container ",err)
		}

	}

	// Delete Switch
	if swCreate {

		// Delete OVS switch
		PullCmds := []string{
			"sudo rm -rf "+ projectPath + "/" + swlogsLocalDirectory,
			"sudo rm -rf "+ projectPath + "/" + swconfLocalDirectory,
			"sudo docker container stop" + " " +
				devName,
			"sudo docker container rm" + " " +
				devName,
		}

		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot stop and delete switch container ",err)
		}

	}

	// Delete controller
	if ctlCreate {

		// Delete Onos controller
		PullCmds := []string{
			"sudo rm -rf "+ projectPath + "/" + ctlogsLocalDirectory,
			"sudo rm -rf "+ projectPath + "/" + ctlconfLocalDirectory,
			"sudo docker container stop" + " " +
				devName,
			"sudo docker container rm" + " " +
				devName,
		}

		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot stop and delete controller container",err)
		}

	}



}