/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"wlicense/util"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists Available Licenses",
	Long:  `Lists all the available licences`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available Licences: ")
		var licenses = util.GetLicensePretty()
		var uglyLicences = util.GetLicenses()

		for j := 1; j <= len(licenses); j++ {
			fmt.Println(strconv.Itoa(j) + ". " + licenses[j-1] + " - " + uglyLicences[j-1])
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
