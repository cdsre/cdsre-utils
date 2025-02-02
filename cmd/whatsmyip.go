/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var ipVersion int

// whatsmyipCmd represents the whatsmyip command
var whatsmyipCmd = &cobra.Command{
	Use:   "whatsmyip",
	Short: "Prints your public IP address",
	Long:  "This util will print your IPv4 address or IPv6 address.",
	Run: func(cmd *cobra.Command, args []string) {
		if ipVersion != 4 && ipVersion != 6 {
			fmt.Printf("Invalid IP version: %d\n", ipVersion)
		}
		fmt.Println(WhatsMyIP())
	},
}

func init() {
	rootCmd.AddCommand(whatsmyipCmd)
	whatsmyipCmd.Flags().IntVarP(&ipVersion, "version", "v", 4, "IP version (4 or 6)")
}

func WhatsMyIP() string {
	res, err := http.Get(fmt.Sprintf("https://api%v.ipify.org", ipVersion))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "Error fetching IP"
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error: %v\n", res.Status)
		return "Error fetching IP"
	}

	ip, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return "Error fetching IP"
	}

	return string(ip)
}
