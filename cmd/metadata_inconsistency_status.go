package cmd

import (
	"fmt"
	"os"

	"github.com/sho-hata/hasuraf/cmd/hasura"
	"github.com/spf13/cobra"
)

var metadataInconsistencyStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check if the metadata is inconsistent or not.",
	Run: func(cmd *cobra.Command, args []string) {
		if result, err := hasura.NewHasuraCmd("metadata inconsistency status", setFlagValues(cmd)).Run(); err != nil {
			fmt.Println(result, err)
			os.Exit(1)
		} else {
			fmt.Println(result)
			os.Exit(0)
		}
	},
}

func init() {
	setFlags(metadataInconsistencyStatusCmd)
}
