package main

// ====================
//  IMPORTS
// ====================

import (
	"log"
	"os"

	"github.com/cosasdepuma/masterchef/backend"
)

// ====================
//  PRIVATE METHODS
// ====================

func main() {
	exitCode := 0
	if err := backend.Serve(); err != nil {
		log.Fatal(err)
		exitCode = 1
	}
	os.Exit(exitCode)
}
