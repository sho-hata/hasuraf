/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/sho-hata/hasura-fzf/cmd/hasura"
	"github.com/spf13/cobra"
)

const use string = "seed"

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   use,
	Short: "Find the seed file to apply and run the \"hasura seed apply\" command.",
	Long: `
Find the seed file to apply and run the \"hasura seed apply'" command.
# It will convert as follows
hasuraf seed -> hasura seed apply --file XXX

# caution
When you use it, put the .env file with "HASURA_GRAPHQL_DATABASE_URL" in the "/hasura" directory.
If the file is located elsewhere, use the \"--envfile\" option to specify the location of the .env file.
`,
	Example: `  hasuraf seed`,
	Run: func(cmd *cobra.Command, args []string) {
		flagOptions := setFlags(cmd)
		if result, err := hasura.NewHasuraCmd(use, flagOptions).Run(); err != nil {
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
	seedCmd.Flags().String("admin-secret", "", "admin secret for Hasura GraphQL engine")
	seedCmd.Flags().String("certificate-authority", "", "path to a cert file for the certificate authority")
	seedCmd.Flags().String("database-name", "default", "database on which operation should be applied (default \"default\")")
	seedCmd.Flags().Bool("disable-interactive", false, "disables interactive prompts (default: false)")
	seedCmd.Flags().String("endpoint", "", "endpoint for Hasura GraphQL engine")
	seedCmd.Flags().String("envfile", ".env", ".env filename to load ENV vars from (default \".env\")")
	seedCmd.Flags().Bool("insecure-skip-tls-verify", false, "skip TLS verification and disable cert checking (default: false) ")
	seedCmd.Flags().String("log-level", "INFO", "log level (DEBUG, INFO, WARN, ERROR, FATAL) (default \"INFO\")")
	seedCmd.Flags().Bool("no-color", false, "do not colorize output (default: false)")
	seedCmd.Flags().String("project", "", "directory where commands are executed (default: current dir)")
	seedCmd.Flags().String("skip-update-check", "", "skip automatic update check on command execution")
}

func setFlags(cmd *cobra.Command) map[string]interface{} {
	flagOptions := map[string]interface{}{}
	if adminSecretFlag, _ := cmd.Flags().GetString("admin-secret"); adminSecretFlag != "" {
		flagOptions["admin-secret"] = adminSecretFlag
	}
	if certificateAuthorityFlag, _ := cmd.Flags().GetString("certificate-authority"); certificateAuthorityFlag != "" {
		flagOptions["certificate-authority"] = certificateAuthorityFlag
	}
	if databaseNameFlag, _ := cmd.Flags().GetString("database-name"); databaseNameFlag != "" {
		flagOptions["database-name"] = databaseNameFlag
	}
	if disableInteractiveFlag, _ := cmd.Flags().GetBool("disable-interactive"); !disableInteractiveFlag {
		flagOptions["disable-interactive"] = disableInteractiveFlag
	}
	if endpointFlag, _ := cmd.Flags().GetString("endpoint"); endpointFlag != "" {
		flagOptions["endpoint"] = endpointFlag
	}
	if envFileFlag, _ := cmd.Flags().GetString("envfile"); envFileFlag != "" {
		flagOptions["envfile"] = envFileFlag
	}
	if insecureSkipTlsVerifyFlag, _ := cmd.Flags().GetBool("insecure-skip-tls-verify"); !insecureSkipTlsVerifyFlag {
		flagOptions["insecure-skip-tls-verify"] = insecureSkipTlsVerifyFlag
	}
	if logLevelFlag, _ := cmd.Flags().GetString("log-level"); logLevelFlag != "" {
		flagOptions["log-level"] = logLevelFlag
	}
	if noColorFlag, _ := cmd.Flags().GetBool("no-color"); !noColorFlag {
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
