// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
)

// adminCmd represents the admin command
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Utilities for helping administration",
	Long: `All command below are tools for administration purposes
    
    Command:
    listUserDeactivated         To get list of Deactivated and Password Change Needed users
    listActiveObj               To get list of activate objects on HANA repository
    listInactiveObj             To get list of inactive objects on HANA repository
    getMemoryConsume            To get HANA Memory Consumption`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("admin called")
		// fmt.Println(cmd.Short)
		// fmt.Println()
		// fmt.Println(cmd.Long)
		cmd.Help()
		// if err := cmd.Execute(); err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }
	},
}

func init() {
	rootCmd.AddCommand(adminCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adminCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// adminCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
