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

var packageNameDPTT, schemaNameDPTT string

// depPkgToTableCmd represents the depPkgToTable command
var depPkgToTableCmd = &cobra.Command{
	Use:   "depPkgToTable",
	Short: "A look up tool to search Table used within specific Package",
	Long: `Look up tool to dependencies between HANA models and specific Table.
	ie. By searching "plantation-slowmove" package, you can get all tables used in those package.
	You also can specify schema name to show only tables from specific schema.
	Information returned are respecting schema name, and also tables.

	Use format: 
	depPkgToTable -p <package_name>
	depPkgToTable -p <package_name> -s <schema_name>`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("depPkgToTable called")

		const (
			depPkgToTableSQL = `SELECT DISTINCT
BASE_SCHEMA_NAME,
BASE_OBJECT_NAME
FROM     "SYS" ."OBJECT_DEPENDENCIES"
WHERE DEPENDENT_SCHEMA_NAME = 'PUBLIC'
AND LEFT(DEPENDENT_OBJECT_NAME, locate(DEPENDENT_OBJECT_NAME, '::') -1) = ?`
			schemaSQL = `AND BASE_SCHEMA_NAME = `
		)

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

		// so manual
		var depPkgToTableSQLComplete string
		if schemaNameDPTT != "" {
			depPkgToTableSQLComplete = depPkgToTableSQL + schemaSQL + "'" + schemaNameDPTT + "'"
		} else {
			depPkgToTableSQLComplete = depPkgToTableSQL
		}

		rows, err := db.Query(depPkgToTableSQLComplete, packageNameDPTT)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// fmt.Println("Parameter:", packageNameDPTT)

		// Write Header
		fmt.Printf("\n|%-20s|%-60s|\n\n", "SCHEMA_NAME", "TABLE_NAME")
		for rows.Next() {
			var schemaName, tableName string
			if err := rows.Scan(&schemaName, &tableName); err != nil {
				// utils.WriteMsg("SCAN")
				log.Fatal(err)
			}
			fmt.Printf("|%-20s|%-60s|\n", schemaName, tableName)
		}
	},
}

func init() {
	devCmd.AddCommand(depPkgToTableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// depPkgToTableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	depPkgToTableCmd.Flags().StringVarP(&packageNameDPTT, "packagename", "p", "", "Package name you want to get dependency (ie. plantation-slowmove)")
	depPkgToTableCmd.MarkFlagRequired("packagename")

	depPkgToTableCmd.Flags().StringVarP(&schemaNameDPTT, "schemaname", "s", "", "Schema name you want to filter")
}
