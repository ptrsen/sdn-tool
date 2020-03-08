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

// patchCmd represents the patch command
var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Add patch between bridges",
	Long: `Add patch between bridges. For example:
			sudo ./sdntool patch --create --switchname ovs1 --bridgeA br1 --bridgeB br2
			sudo ./sdntool patch --delete --switchname ovs1 --bridgeA br1 --bridgeB br2`,
	Run: patch,
}

var patchCreate bool
var patchDelete bool

var brnameA string
var brnameB string

func init() {
	rootCmd.AddCommand(patchCmd)
	patchCmd.PersistentFlags().BoolVar(&patchCreate, "create",false,"Create patch between bridge")
	patchCmd.PersistentFlags().BoolVar(&patchDelete, "delete",false,"Delete patch between bridge")

	patchCmd.Flags().StringVarP(&swname, "switchname", "s", "", "Container switch name")
	patchCmd.Flags().StringVarP(&brnameA, "bridgeA", "a", "", " First bridge ")
	patchCmd.Flags().StringVarP(&brnameB, "bridgeB", "b", "", " Second bridge ")

}

func patch (cmd *cobra.Command, args []string) {

	// Create patch
	if patchCreate{

		PullCmds := []string{
			"sudo docker exec " + swname + " " + "ovs-vsctl --may-exist add-port " + brnameA + " patch" + brnameA + "-" + brnameB ,
			"sudo docker exec " + swname + " " + "ovs-vsctl set interface" + " patch" + brnameA + "-" + brnameB + " type=patch options:peer=" + "patch" + brnameB + "-" + brnameA,
			"sudo docker exec " + swname + " " + "ovs-vsctl --may-exist add-port " + brnameB + " patch" + brnameB + "-" + brnameA ,
			"sudo docker exec " + swname + " " + "ovs-vsctl set interface" + " patch" + brnameB + "-" + brnameA + " type=patch options:peer=" + "patch" + brnameA + "-" + brnameB,
		}

		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot create patch between bridges",err)
		}

	}

	// Delete patch
	if patchDelete {

		PullCmds := []string{
			"sudo docker exec " + swname + " " + "ovs-vsctl --if-exist del-port " + brnameA + " patch" + brnameA + "-" + brnameB,
			"sudo docker exec " + swname + " " + "ovs-vsctl --if-exist del-port " + brnameB + " patch" + brnameB + "-" + brnameA,
		}

		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot delete patch between bridges ",err)
		}

	}

}