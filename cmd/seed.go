package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const useSeed = "seed"

var seedApplyHelpText = `
Manage seed data.

Available Commands:
  apply       Find the file to apply and run the "hasura seed apply" command.`

var seedCmd = &cobra.Command{
	Use:   useSeed,
	Short: "Manage seed data.",
	Long:  seedApplyHelpText,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(seedApplyHelpText)
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
	setFlags(seedCmd)
}
