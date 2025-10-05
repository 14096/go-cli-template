package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

func PrintColor(text string, color string) {
	fmt.Print(color + text + Reset + "\n")
}

func ClearConsole() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Header() {
	header := `  
   ▛▛▌▌▌▀▌▛▌▛▌
   ▌▌▌▙▌█▌▙▌▙▌
      ▄▌  ▌ ▌
	`
	PrintColor(header, Cyan)
	println("\n")
}
