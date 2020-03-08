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
)


// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs necessary things for emulation",
	Long: `Installs necessary things for network emulation. For example:
             sudo ./sdntool install --docker`,
	Run: install,
}

var installDocker bool

func init() {
	rootCmd.AddCommand(installCmd)
	installCmd.PersistentFlags().BoolVar(&installDocker, "docker",false,"Installs Docker CE")
}


func install (cmd *cobra.Command, args []string) {

	if installDocker {

		fmt.Println("Install Step ...")

		// Install docker engine CE
		// https://docs.docker.com/install/linux/docker-ce/ubuntu/

		installCmds := []string{
			"sudo apt update",  // Update enviroment
			"sudo apt -y dist-upgrade",
			"sudo apt -y purge docker docker-engine docker.io containerd runc",   // Remove old docker version
			"sudo apt -y purge docker-ce docker-ce-cli containerd.io",            // Remove new docker version
			"sudo rm -rf /var/lib/docker",
			"sudo apt -y install apt-transport-https ca-certificates curl gnupg-agent software-properties-common bridge-utils net-tools", // Install pre-reqs HTTPS for apt
			"curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -", // Add GPG key for official docker repository
			"sudo apt-key fingerprint 0EBFCD88", //  Check GPC Key
			"sudo add-apt-repository \"deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable\"", // Add docker latest repository to apt
			"sudo apt update",  // Update enviroment
			"sudo apt -y dist-upgrade",
			"apt-cache policy docker-ce", // Check recommended docker source and installation version
			"sudo apt -y install docker-ce docker-ce-cli containerd.io", // Install latest docker version
			"sudo docker --version", // Checking version
			"sudo docker run hello-world", // Verifying Docker
		}

		for _, cmd := range installCmds {
			err := system.ShellExec("/bin","sh", "-c", cmd)
			system.CheckError("Cannot install docker ",err)
		}

		fmt.Println("Docker Installed!")

	}


}

