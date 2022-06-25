package cmd

import (
	"fmt"
	"os"

	"github.com/sho-hata/hasuraf/cmd/hasura"
	"github.com/spf13/cobra"
)

var migrateDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Find the version to delete and run the \"hasura migrate delete\" command.",
	Long: `Find the version to apply and run the "hasura migrate delete" command.
	# It will convert as follows
	hasuraf migrate delete -> hasura migrate delete --version XXX

	# caution
	When you use it, put the .env file with "HASURA_GRAPHQL_DATABASE_URL" in the current directory.
	If the file is located elsewhere, use the "--envfile" option to specify the location of the .env file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if result, err := hasura.NewHasuraCmd(hasura.CalledMigrateDelete, setFlagValues(cmd)).Run(); err != nil {
			fmt.Println(result, err)
			os.Exit(1)
		} else {
			fmt.Println(result)
			os.Exit(0)
		}
	},
}

func init() {
	setFlags(migrateDeleteCmd)
}
