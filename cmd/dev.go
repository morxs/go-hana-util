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

var csvMode bool
var packageName, schemaName string

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Utilities for helping development",
	Long: `All command below are tools for development purposes
    
	Command:
	depTableToModel             A look up tool to search HANA models that using specific Table
	depPkgToTable               A look up tool to search Tables used by specific HANA Package
	listStrucPkgTable           List out structure of relationship HANA Package, View, Schema and Table`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("dev called")
		// fmt.Println(cmd.Short)
		// fmt.Println()
		// fmt.Println(cmd.Long)
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(devCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// devCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// devCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// only for depTableToMode
	depTableToModelCmd.Flags().BoolVarP(&csvMode, "csv", "m", false, "Printout as csv format")

	// only for depPkgToTable
	depPkgToTableCmd.Flags().BoolVarP(&csvMode, "csv", "m", false, "Printout as csv format")

	depPkgToTableCmd.Flags().StringVarP(&packageName, "packagename", "p", "", "Package name you want to get dependency (ie. plantation-slowmove)")
	depPkgToTableCmd.MarkFlagRequired("packagename")
	depPkgToTableCmd.Flags().StringVarP(&schemaName, "schemaname", "s", "", "Schema name you want to filter")

	// only for listStrucPkgTable
	listStrucPkgTableCmd.Flags().StringVarP(&packageName, "packagename", "p", "", "Package name you want to get dependency (ie. plantation-slowmove)")
	listStrucPkgTableCmd.MarkFlagRequired("packagename")
	listStrucPkgTableCmd.Flags().StringVarP(&schemaName, "schemaname", "s", "", "Schema name you want to filter")
}
