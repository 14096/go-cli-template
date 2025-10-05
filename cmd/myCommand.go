package cmd

import (
	"fmt"
	"log/slog"
	"myApp/utils"

	"github.com/spf13/cobra"
)

var myCommandCmd = &cobra.Command{
	Use:    "my-command",
	Short:  "My command",
	PreRun: baseCmd.PreRun,
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		mail, _ := cmd.Flags().GetString("mail")
		myCommand(user, mail)
	},
}

func myCommand(user string, mail string) {
	slog.Info("START\n")

	bar := utils.NewSpinner("my command")
	bar.Add(1)
	bar.Finish()

	fmt.Printf("USER: %s\n", user)
	fmt.Printf("MAIL: %s\n", mail)
	println()

	slog.Info("END\n\n")
}
