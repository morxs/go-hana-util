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

// getUserLoginInfoCmd represents the getUserLoginInfo command
var getUserLoginInfoCmd = &cobra.Command{
	Use:   "getUserLoginInfo",
	Short: "Get users login information",
	Long: `Get user login informations, such as: Validity, Last succesfull and unsuccessful 
connection, Lifetime checks.`,
	Run: func(cmd *cobra.Command, args []string) {
		const (
			getUserLoginInfoSQL = `SELECT 
USER_NAME
, VALID_FROM
, VALID_UNTIL
, LAST_SUCCESSFUL_CONNECT
, LAST_INVALID_CONNECT_ATTEMPT
, IS_PASSWORD_LIFETIME_CHECK_ENABLED
FROM USERS`
		)
		// fmt.Println("getUserLoginInfo called")
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

		rows, err := db.Query(getUserLoginInfoSQL)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Write Header
		fmt.Printf("\n|%-40s|%15s|%15s|%15s|%15s|%20s|\n\n", "USER_NAME", "VALID_FROM", "VALID_UNTIL", "LAST_SUCCESS", "LAST_FAILED", "LIFETIME_CHECK")
		for rows.Next() {
			var userName, lifetimeCheck string
			var validFrom, validUntil, lastSuccess, lastFailed utils.NullTime
			if err := rows.Scan(&userName, &validFrom, &validUntil, &lastSuccess, &lastFailed, &lifetimeCheck); err != nil {
				// utils.WriteMsg("SCAN")
				log.Fatal(err)
			}
			// fmt.Printf("%T %T %T %T %T %T\n", userName, validFrom, validUntil, lastSuccess, lastFailed, lifetimeCheck)
			// if validFrom == nil {
			// 	validFrom = ""
			// } else {
			// 	switch validFrom.(type) {
			// 	case time.Time:
			// 		sValidFrom := time.Time(validFrom).String()
			// 	}
			// }
			// if validUntil == nil {
			// 	validUntil = ""
			// }
			// if lastSuccess == nil {
			// 	lastSuccess = ""
			// }:w

			// if lastFailed == nil {
			// 	lastFailed = ""
			// }
			fmt.Printf("|%-40s|%15s|%15s|%15s|%15s|%20s|\n", userName,
				validFrom.Format("2006-01-02"),
				validUntil.Format("2006-01-02"),
				lastSuccess.Format("2006-01-02"),
				lastFailed.Format("2006-01-02"),
				lifetimeCheck)
		}
	},
}

func init() {
	adminCmd.AddCommand(getUserLoginInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getUserLoginInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getUserLoginInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
