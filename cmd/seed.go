package cmd

import (
	"fmt"
	"os"

	"github.com/sho-hata/hasura-fzf/cmd/hasura"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Find the seed file to apply and run the \"hasura seed apply\" command.",
	Long: `
Find the seed file to apply and run the \"hasura seed apply'" command.
# It will convert as follows
hasuraf seed -> hasura seed apply --file XXX

# caution
When you use it, put the .env file with "HASURA_GRAPHQL_DATABASE_URL" in the "current" directory.
If the file is located elsewhere, use the \"--envfile\" option to specify the location of the .env file.`,
	Example: `
# Apply only a particular file:
hasuraf seed

# Use with admin secret:
hasuraf seed --admin-secret "<admin-secret>"`,
	Run: func(cmd *cobra.Command, args []string) {
		if result, err := hasura.NewHasuraCmd("seed", setFlagValues(cmd)).Run(); err != nil {
			fmt.Println(result, err)
			os.Exit(1)
		} else {
			fmt.Println(result)
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
	setFlags(seedCmd)
}
