package cmd

import (
	"fmt"
	"os"

	"github.com/sho-hata/hasura-fzf/cmd/hasura"
	"github.com/spf13/cobra"
)

var metadataInconsistencyCmd = &cobra.Command{
	Use:     "inconsistency",
	Short:   "Manage inconsistent objects in Hasura metadata.",
	Aliases: []string{"inconsistencies", "ic"},
	Run: func(cmd *cobra.Command, args []string) {
		if result, err := hasura.NewHasuraCmd("metadata inconsistency", setFlagValues(cmd)).Run(); err != nil {
			fmt.Println(result, err)
			os.Exit(1)
		} else {
			fmt.Println(result)
			os.Exit(0)
		}
	},
}

func init() {
	setFlags(metadataInconsistencyCmd)
}
