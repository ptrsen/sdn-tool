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



// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs necessary things from scripts",
	Long: `Installs necessary things include docker`,
	Run: install,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func install (cmd *cobra.Command, args []string) {
	fmt.Println("Install Step ...")


	projectPath, err := os.Getwd()
	system.CheckError([]byte("Cant find Project path"), err)
	projectPath = projectPath[:len(projectPath)-3]

	err = system.ShellExec("/bin","bash", projectPath + "/scripts/docker-install.sh")
	system.CheckError([]byte("Cant execute /scripts/docker-install.sh"),err)

	fmt.Println("Docker Installed!")

}

