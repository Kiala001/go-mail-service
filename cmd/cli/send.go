package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Enviar email",
	Run: func(cmd *cobra.Command, args []string) {
		body := map[string]any{
			"to":      []string{to},
			"subject": subject,
			"html":    html,
			"from":    from,
		}

		data, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
		req.Header.Set("Authorization", "Bearer "+apiKey)
		req.Header.Set("Content-Type", "application/json")

		http.DefaultClient.Do(req)
	},
}

var to, subject, html, from, apiKey, url string

func init() {
	sendCmd.Flags().StringVar(&to, "to", "", "")
	sendCmd.Flags().StringVar(&subject, "subject", "", "")
	sendCmd.Flags().StringVar(&html, "html", "", "")
	sendCmd.Flags().StringVar(&from, "from", "", "")
	sendCmd.Flags().StringVar(&apiKey, "key", "", "")
	sendCmd.Flags().StringVar(&url, "url", "http://localhost:8080/api/send-email", "")
	rootCmd.AddCommand(sendCmd)
}