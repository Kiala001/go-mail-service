package main

import (
	"fmt"

	"github.com/kiala001/go-mail-service/internal/apikey"
	"github.com/spf13/cobra"
)

var apikeyCmd = &cobra.Command{
	Use:   "apikey",
	Short: "Gerenciar API keys",
}

var apikeyGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Gerar nova API key",
	Run: func(cmd *cobra.Command, args []string) {
		key, err := apikey.Generate()
		if err != nil {
			fmt.Println("Erro ao gerar API key:", err)
			return
		}
		fmt.Println(key)
	},
}

func init() {
	apikeyCmd.AddCommand(apikeyGenerateCmd)
	rootCmd.AddCommand(apikeyCmd)
}
