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

var packageNameSPT, schemaNameSPT string

// listStrucPkgTableCmd represents the listStrucPkgTable command
var listStrucPkgTableCmd = &cobra.Command{
	Use:   "listStrucPkgTable",
	Short: "List out structure of relationship HANA Package, View, Schema and Table",
	Long: `List out structure of relationship HANA Package, View, Schema and Table

	Use format: 
	listStrucPkgTable -p <package_name>
	listStrucPkgTable -p <package_name> -s <schema_name>`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("listStrucPkgTable called")

		const (
			topSQL = `SELECT
"PKG_NAME", "MODEL_NAME", "BASE_SCHEMA_NAME", "BASE_OBJECT_NAME"
FROM 
(`
			innerSQL = `SELECT DISTINCT
	LEFT(DEPENDENT_OBJECT_NAME, locate(DEPENDENT_OBJECT_NAME, '::') -1) as "PKG_NAME"
	, RIGHT(DEPENDENT_OBJECT_NAME, length(DEPENDENT_OBJECT_NAME) - locate(DEPENDENT_OBJECT_NAME,'::')-1) as "MODEL_NAME"
	, BASE_SCHEMA_NAME
	, BASE_OBJECT_NAME
   FROM     "SYS" ."OBJECT_DEPENDENCIES"
   WHERE DEPENDENT_SCHEMA_NAME = 'PUBLIC'
   AND LEFT(DEPENDENT_OBJECT_NAME, locate(DEPENDENT_OBJECT_NAME, '::') -1) = ?`
			schemaSQL = `	AND BASE_SCHEMA_NAME = `
			bottomSQL = `)
ORDER BY "PKG_NAME", "MODEL_NAME", BASE_SCHEMA_NAME, BASE_OBJECT_NAME`
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

		// so manual - construct sql
		var depPkgToTableSQLComplete string
		if schemaNameSPT != "" {
			depPkgToTableSQLComplete = topSQL + innerSQL + schemaSQL + "'" + schemaNameSPT + "'" + bottomSQL
		} else {
			depPkgToTableSQLComplete = topSQL + innerSQL + bottomSQL
		}

		rows, err := db.Query(depPkgToTableSQLComplete, packageNameSPT)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// fmt.Println("Parameter:", packageName)

		// Write Header
		fmt.Printf("\n|%-30s|%-50s|%-20s|%-30s|\n\n", "PKG_NAME", "MODEL_NAME", "SCHEMA_NAME", "TABLE_NAME")
		for rows.Next() {
			var pkgName, modelName, schemaName, tableName string
			if err := rows.Scan(&pkgName, &modelName, &schemaName, &tableName); err != nil {
				// utils.WriteMsg("SCAN")
				log.Fatal(err)
			}
			fmt.Printf("|%-30s|%-50s|%-20s|%-30s|\n", pkgName, modelName, schemaName, tableName)
		}
	},
}

func init() {
	devCmd.AddCommand(listStrucPkgTableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listStrucPkgTableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listStrucPkgTableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listStrucPkgTableCmd.Flags().StringVarP(&packageNameSPT, "packagename", "p", "", "Package name you want to get dependency (ie. plantation-slowmove)")
	listStrucPkgTableCmd.MarkFlagRequired("packagename")

	listStrucPkgTableCmd.Flags().StringVarP(&schemaNameSPT, "schemaname", "s", "", "Schema name you want to filter")
}
