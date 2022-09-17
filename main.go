// Package main is the entrypoint for the command line executable.
package main

import (
	"github.com/kmtym1998/hasuraenv/cmd"
	log "github.com/sirupsen/logrus"
)

// main is the entrypoint function
func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
