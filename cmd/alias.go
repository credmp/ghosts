/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// aliasCmd represents the alias command
var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Alias a name to an existing hostname",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("alias called")
	},
}

func init() {
	rootCmd.AddCommand(aliasCmd)
}
