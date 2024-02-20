/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"gitlab.com/credmp/ghosts/internal"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [hostname] [ip]",
	Short: "Add a hostname with an ip address",
	Run:   addHostname,
	Args:  cobra.MinimumNArgs(2),
}

func addHostname(cmd *cobra.Command, args []string) {
	ip := net.ParseIP(args[1])
	if ip == nil {
		log.Fatal("Could not parse IP")
	}

	hf := internal.NewHostFile(hostsfile)

	fmt.Printf("%v\n", hf)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
