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
	"fmt"
	"github.com/ptrsen/sdn-tool/internal/pkg/system"
	"github.com/spf13/cobra"
	"os"
)

// imagesCmd represents the images command
var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Create/Delete docker base images for network devices",
	Long: `Create/Delete docker base images for network devices. For example:
             sudo ./sdntool images --create 
             sudo ./sdntool images --delete`,
	Run: images,
}

var imageCreate bool
var imageDelete bool

func init() {
	rootCmd.AddCommand(imagesCmd)
	imagesCmd.PersistentFlags().BoolVar(&imageCreate, "create",false,"Create all network device images from Dockerfiles")
	imagesCmd.PersistentFlags().BoolVar(&imageDelete, "delete",false,"Delete all network device images from Dockerfiles")
}


func images (cmd *cobra.Command, args []string) {

	projectPath, er := os.Getwd()
	system.CheckError("Cant find Project path", er)

	//  Pull Docker base images
	if imageCreate {

		controllerimage := projectPath +"/dockerfiles/controller"
		switchimage := projectPath +"/dockerfiles/switch"
		hostimage := projectPath +"/dockerfiles/host"

		PullCmds := []string{
			"sudo docker build -t controller "+ controllerimage,
			"sudo docker build -t switch "+ switchimage,
			"sudo docker build -t host "+ hostimage,
		}

		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot pull images from dockerfiles ",err)
		}
		fmt.Println("All images created !!")
	}

	//  Prune Docker base images
	if imageDelete {

		PullCmds := []string{
			"yes | sudo docker image prune -a",
		}

		for _, cmd := range PullCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot delete images ",err)
		}
		fmt.Println("All images deleted !!")

	}


}