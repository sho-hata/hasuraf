package cmd

import (
	"fmt"
	"os"

	"github.com/sho-hata/hasura-fzf/cmd/hasura"
	"github.com/spf13/cobra"
)

const useMetadata = "metadata"

var metadataCmd = &cobra.Command{
	Use:     useMetadata,
	Aliases: []string{"md"},
	Short:   "Manage Hasura GraphQL engine metadata saved in the database",
	Run: func(cmd *cobra.Command, args []string) {
		if result, err := hasura.NewHasuraCmd("metadata", setFlagValues(cmd)).Run(); err != nil {
			fmt.Println(result, err)
			os.Exit(1)
		} else {
			fmt.Println(result)
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.AddCommand(metadataCmd)
	metadataCmd.AddCommand(metadataApplyCmd)
	setFlags(metadataCmd)
}
