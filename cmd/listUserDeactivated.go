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

	// Register hdb driver.
	_ "github.com/SAP/go-hdb/driver"
)

// listUserDeactivatedCmd represents the listUserDeactivated command
var listUserDeactivatedCmd = &cobra.Command{
	Use:   "listUserDeactivated",
	Short: "List Deactivated Users",
	Long:  `List all deactivated users, maybe due to incorrect password attempt or locked`,
	Run: func(cmd *cobra.Command, args []string) {
		const (
			listUserDeactivatedSQL = `SELECT
USER_NAME, 
PASSWORD_CHANGE_NEEDED, 
USER_DEACTIVATED
FROM "SYS"."USERS"
WHERE ( USER_DEACTIVATED='TRUE' 
OR PASSWORD_CHANGE_NEEDED='TRUE')`
		)
		// fmt.Println("listUserDeactivated called")
		// fmt.Println("Config file: ", sCfg)
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

		rows, err := db.Query(listUserDeactivatedSQL)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Write Header
		fmt.Printf("\n|%-50s|%-15s|%-15s|\n\n", "USERNAME", "CHG PWD NEED", "DEACTIVATED")
		// fmt.Println("USERNAME, PASSWORD_CHANGE_NEEDED, USER_DEACTIVATED")
		for rows.Next() {
			var userName, passwordChangeNeeded, userDeactivated string
			if err := rows.Scan(&userName, &passwordChangeNeeded, &userDeactivated); err != nil {
				// utils.WriteMsg("SCAN")
				log.Fatal(err)
			}
			fmt.Printf("|%-50s|%-15s|%-15s|\n", userName, passwordChangeNeeded, userDeactivated)
		}
	},
}

func init() {
	adminCmd.AddCommand(listUserDeactivatedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listUserDeactivatedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listUserDeactivatedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
