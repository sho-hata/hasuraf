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
		flagOptions := setFlags(cmd)
		if result, err := hasura.NewHasuraCmd(use, flagOptions).Run(); err != nil {
			log.Fatal(result, err)
		} else {
			fmt.Println(result)
		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
	seedCmd.Flags().String("admin-secret", "", "admin secret for Hasura GraphQL engine")
	seedCmd.Flags().String("certificate-authority", "", "path to a cert file for the certificate authority")
	seedCmd.Flags().String("database-name", "", "database on which operation should be applied")
	seedCmd.Flags().String("disable-interactive", "", "disables interactive prompts (default: false)")
	seedCmd.Flags().String("endpoint", "", "endpoint for Hasura GraphQL engine")
	seedCmd.Flags().String("envfile", "", " .env filename to load ENV vars from (default \".env\")")
	seedCmd.Flags().String("insecure-skip-tls-verify", "", "skip TLS verification and disable cert checking (default: false) ")
	seedCmd.Flags().String("log-level", "", "log level (DEBUG, INFO, WARN, ERROR, FATAL) (default \"INFO\")")
	seedCmd.Flags().String("no-color", "", "do not colorize output (default: false)")
	seedCmd.Flags().String("project", "", "directory where commands are executed (default: current dir)")
	seedCmd.Flags().String("skip-update-check", "", "skip automatic update check on command execution")
}

func setFlags(cmd *cobra.Command) map[string]string {
	flagOptions := map[string]string{}
	if adminSecretFlag, _ := cmd.Flags().GetString("admin-secret"); adminSecretFlag != "" {
		flagOptions["admin-secret"] = adminSecretFlag
	}
	if certificateAuthorityFlag, _ := cmd.Flags().GetString("certificate-authority"); certificateAuthorityFlag != "" {
		flagOptions["certificate-authority"] = certificateAuthorityFlag
	}
	if databaseNameFlag, _ := cmd.Flags().GetString("database-name"); databaseNameFlag != "" {
		flagOptions["database-name"] = databaseNameFlag
	}
	if disableInteractiveFlag, _ := cmd.Flags().GetString("disable-interactive"); disableInteractiveFlag != "" {
		flagOptions["disable-interactive"] = disableInteractiveFlag
	}
	if endpointFlag, _ := cmd.Flags().GetString("endpoint"); endpointFlag != "" {
		flagOptions["endpoint"] = endpointFlag
	}
	if envFileFlag, _ := cmd.Flags().GetString("envfile"); envFileFlag != "" {
		flagOptions["envfile"] = envFileFlag
	}
	if insecureSkipTlsVerifyFlag, _ := cmd.Flags().GetString("insecure-skip-tls-verify"); insecureSkipTlsVerifyFlag != "" {
		flagOptions["insecure-skip-tls-verify"] = insecureSkipTlsVerifyFlag
	}
	if logLevelFlag, _ := cmd.Flags().GetString("log-level"); logLevelFlag != "" {
		flagOptions["log-level"] = logLevelFlag
	}
	if noColorFlag, _ := cmd.Flags().GetString("no-color"); noColorFlag != "" {
		flagOptions["no-color"] = noColorFlag
	}
	if projectFlag, _ := cmd.Flags().GetString("project"); projectFlag != "" {
		flagOptions["project"] = projectFlag
	}
	if skipUpdateCheck, _ := cmd.Flags().GetString("skip-update-check "); skipUpdateCheck != "" {
		flagOptions["skip-update-check "] = skipUpdateCheck
	}
	return flagOptions
}
