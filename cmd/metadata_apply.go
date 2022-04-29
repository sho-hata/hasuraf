package cmd

import (
	"fmt"
	"os"

	"github.com/sho-hata/hasuraf/cmd/hasura"
	"github.com/spf13/cobra"
)

var metadataApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply Hasura metadata on a database.",
	Run: func(cmd *cobra.Command, args []string) {
		if result, err := hasura.NewHasuraCmd("metadata apply", setFlagValues(cmd)).Run(); err != nil {
			fmt.Println(result, err)
			os.Exit(1)
		} else {
			fmt.Println(result)
			os.Exit(0)
		}
	},
}

func init() {
	setFlags(metadataApplyCmd)
}
