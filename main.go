package main

import (
	"go-commander/cmd"

	"github.com/charmbracelet/log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		// Handle any errors returned by the root command execution
		log.Fatal(err)
	}
}
