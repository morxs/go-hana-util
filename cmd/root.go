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
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// var cfgFile string
var sCfg string

// var Region string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-hana-util",
	Short: "A collection of HANA utility",
	Long: `So tired to administer HANA DB? I am too. That's reason to create this utility.
	
	Command:
	admin
	dev`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
	rootCmd.PersistentFlags().StringVarP(&sCfg, "config", "c", "", "Config filename (default: config.ini)")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func init() {
// 	cobra.OnInitialize(initConfig)
// 	fmt.Println(sCfg)
// 	// Here you will define your flags and configuration settings.
// 	// Cobra supports persistent flags, which, if defined here,
// 	// will be global for your application.
// 	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-hana-util.yaml)")
// 	fmt.Println(sCfg)
// 	// Cobra also supports local flags, which will only run
// 	// when this action is called directly.
// 	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

// func initConfig() {
// 	utils.WriteMsg("READ CONFIG")
// 	if sCfg != "" {
// 		sCfg = cfgFile
// 	} else {
// 		sCfg = "config.ini"
// 	}
// 	hdbDsn, err := utils.ReadConfig(sCfg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 		viper.SetConfigType("json")
// 	} else {
// 		// Find home directory.
// 		home, err := homedir.Dir()
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		// Search config in home directory with name ".go-hana-util" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".go-hana-util")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
