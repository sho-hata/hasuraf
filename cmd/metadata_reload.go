package cmd

import (
	"fmt"
	"os"

	"github.com/sho-hata/hasuraf/cmd/hasura"
	"github.com/spf13/cobra"
)

var metadataReloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload Hasura GraphQL engine metadata on the database.",
	Run: func(cmd *cobra.Command, args []string) {
		if result, err := hasura.NewHasuraCmd("metadata reload", setFlagValues(cmd)).Run(); err != nil {
			fmt.Println(result, err)
			os.Exit(1)
		} else {
			fmt.Println(result)
			os.Exit(0)
		}
	},
}

func init() {
	setFlags(metadataReloadCmd)
}
