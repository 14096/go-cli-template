package cmd

import (
	"flag"
	"fmt"
	"os"
)

func runMyCommand(args []string) {
	fs := flag.NewFlagSet("my-command", flag.ExitOnError)
	var user, email string
	fs.StringVar(&user, "user", "", "Description for user")
	fs.StringVar(&user, "u", "", "Description for user")
	fs.StringVar(&email, "email", "", "Description for email")
	fs.StringVar(&email, "e", "", "Description for email")
	fs.Parse(args)

	if user == "" {
		fmt.Fprintln(os.Stderr, "error: required flag \"user\" not set")
		fs.Usage()
		os.Exit(1)
	}
	if email == "" {
		fmt.Fprintln(os.Stderr, "error: required flag \"email\" not set")
		fs.Usage()
		os.Exit(1)
	}

	myCommand(user, email)
}

func myCommand(user string, email string) {
	fmt.Printf("USER: %s\n", user)
	fmt.Printf("MAIL: %s\n", email)

	println("\n")
}
