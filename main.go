package main

import (
	"os"

	"github.com/artemskriabin/email-subscription-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
