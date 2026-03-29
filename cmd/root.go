package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"myApp/utils"

	"github.com/phplego/tint"
)

func setup() {
	utils.ClearConsole()
	utils.Header()

	w := os.Stderr
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.RFC3339,
		}),
	))
}

func printHelp() {
	fmt.Println("myApp command line application")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  myApp <command> [flags]")
	fmt.Println()
	fmt.Println("Available Commands:")
	fmt.Println("  my-command   run my-command")
	fmt.Println()
	fmt.Println("Flags for new-command:")
	fmt.Println("  -u, -user string    (required)")
	fmt.Println("  -e, -email string   (required)")
	println()
}

func Execute() {
	start := time.Now()
	setup()
	defer func() {
		utils.PrintElapsed(time.Since(start))
	}()

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	subcommand := os.Args[1]
	args := os.Args[2:]

	switch subcommand {
	case "my-command":
		runMyCommand(args)
	default:
		printHelp()
	}
}
