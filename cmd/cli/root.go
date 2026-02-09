package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "emailcli",
	Short: "CLI para envio de emails",
}

func Execute() {
	rootCmd.Execute()
}