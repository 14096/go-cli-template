package cmd

import (
	"log/slog"
	"myApp/utils"
	"os"
	"time"

	"github.com/phplego/tint"
	"github.com/spf13/cobra"
)

var baseCmd = &cobra.Command{
	PreRun: func(cmd *cobra.Command, args []string) {
		utils.ClearConsole()
		utils.Header()

		w := os.Stderr
		slog.SetDefault(slog.New(
			tint.NewHandler(w, &tint.Options{
				Level:      slog.LevelDebug,
				TimeFormat: time.RFC3339,
			}),
		))
	},
}

var rootCmd = &cobra.Command{
	Use:    "myApp",
	Short:  "Command line application",
	PreRun: baseCmd.PreRun,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	user string
	mail string
)

func init() {
	rootCmd.AddCommand(myCommandCmd)

	myCommandCmd.Flags().StringVarP(&user, "user", "u", "", "Description for user")
	myCommandCmd.Flags().StringVarP(&mail, "mail", "m", "", "Description for mail")
}
