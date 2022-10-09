// Package main is the entrypoint for the command line executable.
package main

import (
	"github.com/kmtym1998/hasuraenv/commands"
	log "github.com/sirupsen/logrus"
)

// main is the entrypoint function
func main() {
	if err := commands.Execute(); err != nil {
		log.Fatal(err)
	}
}
