package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:  "list",
	Long: `サブコマンド一覧を出力`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, c := range rootCmd.Commands() {
			fmt.Println(c.Use)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
