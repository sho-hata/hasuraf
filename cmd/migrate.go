package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateHelpText = `
Manage migrations on the database.

Available Commands:
  apply       Find the version to apply and run the "hasura migrate apply" command.`

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Manage migrations on the database.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(migrateHelpText)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
