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

const (
	getMemoryConsumeSQL = `SELECT TOP 1000
	TABLE_NAME,
	MEMORY_SIZE_IN_TOTAL,
	MEMORY_SIZE_IN_MAIN,
	MEMORY_SIZE_IN_DELTA,
	ESTIMATED_MAX_MEMORY_SIZE_IN_TOTAL,
	RECORD_COUNT
	FROM M_CS_TABLES
	WHERE SCHEMA_NAME = 'SAPABAP1'
	AND TABLE_NAME not like '/%'
	ORDER BY MEMORY_SIZE_IN_TOTAL DESC`
)

// getMemoryConsumeCmd represents the getMemoryConsume command
var getMemoryConsumeCmd = &cobra.Command{
	Use:   "getMemoryConsume",
	Short: "Get HANA Memory Consumption",
	Long:  `Get full memory overview, delta and consumption to monitor HANA resource`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("getMemoryConsume called")
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

		rows, err := db.Query(getMemoryConsumeSQL)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Write Header
		fmt.Printf("\n|%-40s|%15s|%15s|%15s|%15s|%15s|\n\n", "TABLE_NAME", "MEMORY_TOTAL", "MEMORY_MAIN", "MEMORY_DELTA", "EST_MAX_MEMORY", "RECORD_COUNT")
		for rows.Next() {
			var tableName, memoryTotal, memoryMain, memoryDelta, estMaxMemory, recordCount string
			if err := rows.Scan(&tableName, &memoryTotal, &memoryMain, &memoryDelta, &estMaxMemory, &recordCount); err != nil {
				// utils.WriteMsg("SCAN")
				log.Fatal(err)
			}
			fmt.Printf("|%-40s|%15s|%15s|%15s|%15s|%15s|\n", tableName, memoryTotal, memoryMain, memoryDelta, estMaxMemory, recordCount)
		}
	},
}

func init() {
	adminCmd.AddCommand(getMemoryConsumeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getMemoryConsumeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getMemoryConsumeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
