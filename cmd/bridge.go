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

// bridgeCmd represents the bridge command
var bridgeCmd = &cobra.Command{
	Use:   "bridge",
	Short: "Create/Delete bridge for container switch",
	Long: `Create/Delete bridge for container switch. For example:
			sudo ./sdntool bridge --create --switchname ovs1 --bridgename br1 --controllerip tcp:172.17.0.2:6633
			sudo ./sdntool bridge --delete --switchname ovs1 --bridgename br1`,
	Run: bridge,
}

var bridgeCreate bool
var bridgeDelete bool

var swname string
var brname string
var ctlip string

func init() {
	rootCmd.AddCommand(bridgeCmd)
	bridgeCmd.PersistentFlags().BoolVar(&bridgeCreate, "create",false,"Create bridge for container switch")
	bridgeCmd.PersistentFlags().BoolVar(&bridgeDelete, "delete",false,"Delete bridge for container switch")

	bridgeCmd.Flags().StringVarP(&swname, "switchname", "s", "", "Container switch name")
	bridgeCmd.Flags().StringVarP(&brname, "bridgename", "b", "", " Name for new bridge")
	bridgeCmd.Flags().StringVarP(&ctlip, "controllerip", "c", "", "Controller address for new bridge")
}



func bridge (cmd *cobra.Command, args []string) {

	// Create bridge
	if bridgeCreate{

		// Create OVS bridge
		PullCmds := []string{
			"sudo docker exec " + swname + " " + "ovs-vsctl --may-exist add-br " + brname,
			"sudo docker exec " + swname + " " + "ovs-vsctl set-controller " + brname + " " + ctlip,
		}

		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot create bridge on container switch",err)
		}

	}

	// Delete bridge
	if bridgeDelete {

		PullCmds := []string{
			"sudo docker exec " + swname + " " + "ovs-vsctl --if-exists del-br " + brname,
		}

		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot delete bridge on container switch",err)
		}

	}


}