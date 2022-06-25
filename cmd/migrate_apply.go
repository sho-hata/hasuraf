package cmd

import (
	"fmt"
	"os"

	"github.com/sho-hata/hasuraf/cmd/hasura"
	"github.com/spf13/cobra"
)

var migrateApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Find the version to apply and run the \"hasura migrate apply\" command.",
	Long: `
Find the version to apply and run the "hasura migrate apply" command.
# It will convert as follows
hasuraf migrate apply -> hasura migrate apply --version XXX

# caution
When you use it, put the .env file with "HASURA_GRAPHQL_DATABASE_URL" in the current directory.
If the file is located elsewhere, use the "--envfile" option to specify the location of the .env file.`,
	Example: `
# Apply a particular migration version only:
hasuraf migrate apply

# Use with admin secret:
hasuraf migrate apply --admin-secret "<admin-secret>"`,
	Run: func(cmd *cobra.Command, args []string) {
		if result, err := hasura.NewHasuraCmd(hasura.CalledMigrateApply, setFlagValues(cmd)).Run(); err != nil {
			fmt.Println(result, err)
			os.Exit(1)
		} else {
			fmt.Println(result)
			os.Exit(0)
		}
	},
}

func init() {
	setFlags(migrateApplyCmd)
}
