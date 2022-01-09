/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/sho-hata/hasura-fzf/cmd/hasura"
	"github.com/spf13/cobra"
)

const use string = "seed"

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   use,
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		flagOptions := map[string]string{}
		if databaseNameFlag, _ := cmd.Flags().GetString("database-name"); databaseNameFlag != "" {
			flagOptions["database-name"] = databaseNameFlag
		}
		if adminSecretFlag, _ := cmd.Flags().GetString("admin-secret"); adminSecretFlag != "" {
			flagOptions["admin-secret"] = adminSecretFlag
		}

		if result, err := hasura.NewHasuraCmd(use, flagOptions).Run(); err != nil {
			log.Fatal(result, err)
		} else {
			fmt.Println(result)
		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
	seedCmd.Flags().String("database-name", "", "database on which operation should be applied")
	seedCmd.Flags().String("admin-secret", "", "admin secret for Hasura GraphQL engine")
}
