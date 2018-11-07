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
	"database/sql"
	"fmt"
	"log"

	"github.com/morxs/go-hana/utils"
	"github.com/spf13/cobra"
)

// listInactiveObjCmd represents the listInactiveObj command
var listInactiveObjCmd = &cobra.Command{
	Use:   "listInactiveObj",
	Short: "List activate objects on HANA repository",
	Long: `To list out all inactive objects on HANA repositories. 
	This can come in handy when you need to troubleshoot transport problem on HANA 2.0`,
	Run: func(cmd *cobra.Command, args []string) {
		const (
			listInactiveObjSQL = `SELECT PACKAGE_ID, OBJECT_NAME, OBJECT_SUFFIX FROM "_SYS_REPO"."INACTIVE_OBJECT"
WHERE "OBJECT_SUFFIX" = 'attributeview'
OR "OBJECT_SUFFIX" = 'analyticview'
OR "OBJECT_SUFFIX" = 'calculationview'
ORDER BY PACKAGE_ID`
		) // fmt.Println("listInactiveObj called")
		hdbDsn, err := utils.ReadConfig(sCfg)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(hdbDsn)

		db, err := sql.Open(utils.DriverName, hdbDsn)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}

		rows, err := db.Query(listInactiveObjSQL)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Write Header
		fmt.Printf("\n|%-40s|%-40s|%-20s|\n\n", "PACKAGE ID", "OBJECT_NAME", "OBJECT_TYPE")
		for rows.Next() {
			var packageID, objectName, objectType string
			if err := rows.Scan(&packageID, &objectName, &objectType); err != nil {
				// utils.WriteMsg("SCAN")
				log.Fatal(err)
			}
			fmt.Printf("|%-40s|%-40s|%-20s|\n", packageID, objectName, objectType)
		}
	},
}

func init() {
	adminCmd.AddCommand(listInactiveObjCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listInactiveObjCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listInactiveObjCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
