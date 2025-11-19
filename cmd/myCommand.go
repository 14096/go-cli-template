package cmd

import (
	"fmt"
	"log/slog"

	"github.com/slok/gospinner"
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
	s, _ := gospinner.NewSpinner(gospinner.Dots)
	s.Start("my command...")
	slog.Info("START\n")

	fmt.Printf("USER: %s\n", user)
	fmt.Printf("MAIL: %s\n", mail)

	println()

	slog.Info("END")
	s.FinishWithMessage("✅", "DONE\n\n")
}
