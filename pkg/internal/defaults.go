package internal

// ====================
//  IMPORTS
// ====================

import (
	"log"
	"os"
	"strconv"
)

// ====================
//  GLOBALS
// ====================

var (
	DefaultHost    = "[::1]" // Ipv6
	DefaultPort    = 7767    // Decimal ASCII: MC
	DefaultThreads = 100     // Max number of threads
)

// ====================
//  PUBLIC METHODS
// ====================

func GetEnvironmentConfig() {
	// MCHOST
	if env := os.Getenv("MCHOST"); len(env) > 0 {
		log.Printf("|*| Environment variable: MCHOST=%s\n", env)
		DefaultHost = env
	}
	// MCPORT
	if env := os.Getenv("MCPORT"); len(env) > 0 {
		if port, err := strconv.Atoi(env); err == nil {
			log.Printf("|*| Environment variable: MCPORT=%d\n", port)
			DefaultPort = port
		}
	}
	// MCTHREADS
	if env := os.Getenv("MCTHREADS"); len(env) > 0 {
		if threads, err := strconv.Atoi(env); err == nil {
			log.Printf("|*| Environment variable: MCTHREADS=%d\n", threads)
			DefaultThreads = threads
		}
	}
}
