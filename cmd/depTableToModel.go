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
	depTableToModelObjSQL = `SELECT DISTINCT
LEFT(DEPENDENT_OBJECT_NAME, locate(DEPENDENT_OBJECT_NAME,'::' ) -1) as "Package Name"
, RIGHT(DEPENDENT_OBJECT_NAME, length(DEPENDENT_OBJECT_NAME) - locate(DEPENDENT_OBJECT_NAME,'::')-1) as "Model Name"
FROM     "SYS" ."OBJECT_DEPENDENCIES"
WHERE BASE_OBJECT_NAME = ?
AND DEPENDENT_SCHEMA_NAME = 'PUBLIC'
AND LEFT (DEPENDENT_OBJECT_NAME, LOCATE(DEPENDENT_OBJECT_NAME, '::') -1) <> ''`
)

var tableName string

// depTableToModelCmd represents the depTableToModel command
var depTableToModelCmd = &cobra.Command{
	Use:   "depTableToModel",
	Short: "A look up tool to search HANA models that using specific Table",
	Long: `Look up tool to dependencies between HANA models and specific Table.
	ie. By searching MSEG table, you can get all the HANA models and package that are using those table.
	Information returned are respecting HANA Package, and also HANA models.

	Use format: 
	depTableToModel -t <table_name>`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("depTableToModel called")
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

		rows, err := db.Query(depTableToModelObjSQL, tableName)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Write Header
		fmt.Printf("\n|%-40s|%-60s|\n\n", "PACKAGE", "MODEL_VIEW")
		for rows.Next() {
			var packageName, modelName string
			if err := rows.Scan(&packageName, &modelName); err != nil {
				// utils.WriteMsg("SCAN")
				log.Fatal(err)
			}
			fmt.Printf("|%-40s|%-60s|\n", packageName, modelName)
		}
	},
}

func init() {
	devCmd.AddCommand(depTableToModelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// depTableToModelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	depTableToModelCmd.Flags().StringVarP(&tableName, "tablename", "t", "", "Table name you want to get dependency (ie. MSEG, BSEG)")
	depTableToModelCmd.MarkFlagRequired("tablename")
}
