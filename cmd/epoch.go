/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var epoch int64

// epochCmd represents the epoch command
var epochCmd = &cobra.Command{
	Use:   "epoch",
	Short: "given an epoch, return the date in human readable format",
	Long: `This command will take an epoch and return the date in human readable format.
			Example:
			cdsre-utils epoch 1620000000
`,
	ValidArgs: []string{"epoch"},
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}

		if _, err := strconv.Atoi(args[0]); err != nil {
			return fmt.Errorf("invalid epoch: %v", args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		epoch, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Println(epochToHumanReadable(epoch))
	},
}

func epochToHumanReadable(epoch int64) string {
	t := time.Unix(epoch, 0).UTC()
	return t.Format(time.RFC3339)
}

func init() {
	rootCmd.AddCommand(epochCmd)
	epochCmd.DisableFlagsInUseLine = true
}
