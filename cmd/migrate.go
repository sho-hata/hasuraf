package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const useMigrate = "migrate"

var migrateHelpText = `
Manage migrations on the database.

Available Commands:
  apply       Find the version to apply and run the "hasura migrate apply" command.
  delete      Find the version to delete and run the "hasura migrate delete" command.`

var migrateCmd = &cobra.Command{
	Use:   useMigrate,
	Short: "Manage migrations on the database.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(migrateHelpText)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(migrateApplyCmd, migrateDeleteCmd, migrateStatusCmd)
}
