/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"gitlab.com/credmp/gohed/internal"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [hostname] [ip]",
	Short: "Add a hostname",
	Long:  `Add a hostname to the hosts file`,
	Run:   addHostname,
	Args:  cobra.MinimumNArgs(2),
}

func addHostname(cmd *cobra.Command, args []string) {
	ip := net.ParseIP(args[1])
	if ip == nil {
		log.Fatal("Could not parse IP")
	}
	he := internal.NewHostEntry(args[0], ip)
	fmt.Printf("%v\n", he)
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")

}
