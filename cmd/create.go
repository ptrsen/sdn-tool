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
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)


var numberOVS int

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create OVSs",
	Long: `Create OVSs`,
	//Args:  check,
	Run: create,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().IntVarP(&numberOVS, "ovs", "o", 1, "Number of OVSs")
}


func check (cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires one argument")
	}
	//if myapp.IsValidColor(args[0]) {
	//	return nil
	//}
	return fmt.Errorf("invalid  arg specified: %s", args[0])
}

func create (cmd *cobra.Command, args []string) {
	fmt.Println("Create Step ...")
	fmt.Println("Number of OVS : ", numberOVS)
}