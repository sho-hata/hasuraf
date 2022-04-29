package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const hasurafASCIIText = `
___  ___  ________  ________  ___  ___  ________  ________  ________ 
|\  \|\  \|\   __  \|\   ____\|\  \|\  \|\   __  \|\   __  \|\  _____\
\ \  \\\  \ \  \|\  \ \  \___|\ \  \\\  \ \  \|\  \ \  \|\  \ \  \__/ 
 \ \   __  \ \   __  \ \_____  \ \  \\\  \ \   _  _\ \   __  \ \   __\
  \ \  \ \  \ \  \ \  \|____|\  \ \  \\\  \ \  \\  \\ \  \ \  \ \  \_|
   \ \__\ \__\ \__\ \__\____\_\  \ \_______\ \__\\ _\\ \__\ \__\ \__\ 
    \|__|\|__|\|__|\|__|\_________\|_______|\|__|\|__|\|__|\|__|\|__| 
                       \|_________|                                   

docs: https://github.com/sho-hata/hasuraf
`

var rootCmd = &cobra.Command{
	Use:   "hasuraf",
	Short: "This command has a fzf-like UI that allows you to find and run the file version used by the hasura command.",
	Long:  hasurafASCIIText,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
