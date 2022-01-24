package cmd

import "github.com/spf13/cobra"

func setFlags(cmd *cobra.Command) {
	switch cmd.Use {
	case "migrate", "apply", "delete":
		setMigrateFlags(cmd)
		fallthrough
	case "metadata":
		setOptionalFlags(cmd)
	}

	setGlobalFlags(cmd)
}

func setGlobalFlags(cmd *cobra.Command) {
	cmd.Flags().String("envfile", ".env", ".env filename to load ENV vars from (default \".env\")")
	cmd.Flags().String("log-level", "INFO", "log level (DEBUG, INFO, WARN, ERROR, FATAL) (default \"INFO\")")
	cmd.Flags().Bool("no-color", false, "do not colorize output (default: false)")
	cmd.Flags().String("project", "", "directory where commands are executed (default: current dir)")
	cmd.Flags().String("skip-update-check", "", "skip automatic update check on command execution")
}

func setOptionalFlags(cmd *cobra.Command) {
	cmd.Flags().String("admin-secret", "", "admin secret for Hasura GraphQL engine")
	cmd.Flags().String("certificate-authority", "", "path to a cert file for the certificate authority")
	cmd.Flags().String("endpoint", "", "endpoint for Hasura GraphQL engine")
	cmd.Flags().Bool("insecure-skip-tls-verify", false, "skip TLS verification and disable cert checking (default: false) ")
}

func setMigrateFlags(cmd *cobra.Command) {
	cmd.Flags().String("database-name", "default", "database on which operation should be applied (default \"default\")")
	cmd.Flags().Bool("disable-interactive", false, "disables interactive prompts (default: false)")
}

func setFlagValues(cmd *cobra.Command) map[string]interface{} {
	flagOptions := map[string]interface{}{}

	switch cmd.Use {
	case "migrate", "apply", "delete":
		setMigrateFlagValues(cmd, flagOptions)
		fallthrough
	case "metadata":
		setOptionalFlagValues(cmd, flagOptions)
	}

	setGlobalFlagValues(cmd, flagOptions)
	return flagOptions
}

func setGlobalFlagValues(cmd *cobra.Command, flagOptions map[string]interface{}) {
	if envFileFlag, _ := cmd.Flags().GetString("envfile"); envFileFlag != "" {
		flagOptions["envfile"] = envFileFlag
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
	if skipUpdateCheck, _ := cmd.Flags().GetString("skip-update-check"); skipUpdateCheck != "" {
		flagOptions["skip-update-check"] = skipUpdateCheck
	}
}

func setOptionalFlagValues(cmd *cobra.Command, flagOptions map[string]interface{}) {
	if adminSecretFlag, _ := cmd.Flags().GetString("admin-secret"); adminSecretFlag != "" {
		flagOptions["admin-secret"] = adminSecretFlag
	}
	if certificateAuthorityFlag, _ := cmd.Flags().GetString("certificate-authority"); certificateAuthorityFlag != "" {
		flagOptions["certificate-authority"] = certificateAuthorityFlag
	}
	if endpointFlag, _ := cmd.Flags().GetString("endpoint"); endpointFlag != "" {
		flagOptions["endpoint"] = endpointFlag
	}
	if insecureSkipTlsVerifyFlag, _ := cmd.Flags().GetBool("insecure-skip-tls-verify"); !insecureSkipTlsVerifyFlag {
		flagOptions["insecure-skip-tls-verify"] = insecureSkipTlsVerifyFlag
	}
}

func setMigrateFlagValues(cmd *cobra.Command, flagOptions map[string]interface{}) {
	if databaseNameFlag, _ := cmd.Flags().GetString("database-name"); databaseNameFlag != "" {
		flagOptions["database-name"] = databaseNameFlag
	}
	if disableInteractiveFlag, _ := cmd.Flags().GetBool("disable-interactive"); !disableInteractiveFlag {
		flagOptions["disable-interactive"] = disableInteractiveFlag
	}
}
