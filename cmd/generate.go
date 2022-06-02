/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	s "strings"
	"time"
	"wlicense/util"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates License",
	Long:  `Generates a license for your project`,
	Run: func(cmd *cobra.Command, args []string) {
		license, _ := cmd.Flags().GetString("license")
		year, _ := cmd.Flags().GetInt("year")
		name, _ := cmd.Flags().GetString("name")
		makeFile, _ := cmd.Flags().GetBool("output")

		if util.StringInSlice(license, util.GetLicenses()) {
			licenseText := util.GetLicense(license)
			licenseText = s.Replace(licenseText, "[y]", strconv.Itoa(year), -1)
			licenseText = s.Replace(licenseText, "[ch]", name, -1)
			if makeFile {
				if !util.DoesFileExist("LICENSE") {
					fmt.Println("Create License Logic")
				} else {
					fmt.Println("File already exists!")
				}
			}
			fmt.Println(licenseText)
		} else {
			fmt.Println("License not found!")
		}

		fmt.Println(util.GetLicense("waffle"))
		// fmt.Println(cmd.Flags().GetBool("output"))
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	year, _, _ := time.Now().Date()

	generateCmd.Flags().BoolP("output", "o", false, "Decides whether it will output to file")
	generateCmd.Flags().StringP("license", "l", "mit", "Decides which license your using")
	generateCmd.Flags().IntP("year", "y", year, "Decides what year to insert into the license")
	generateCmd.Flags().StringP("name", "n", "Person", "Decides what name to insert into the license")
}
