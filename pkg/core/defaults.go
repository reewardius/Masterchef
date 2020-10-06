package core

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
	DefaultHost = "[::1]" // Ipv6
	DefaultPort = 7767    // Decimal ASCII: MC
	DefaultChef = ""
)

// ====================
//  PUBLIC METHODS
// ====================

func GetEnvironmentConfig() {
	// MCHOST
	if ehost := os.Getenv("MCHOST"); len(ehost) > 0 {
		log.Printf("|*| Environment variable: MCHOST=%s\n", ehost)
		DefaultHost = ehost
	}
	// MCPORT
	if eport := os.Getenv("MCPORT"); len(eport) > 0 {
		if port, err := strconv.Atoi(eport); err == nil {
			log.Printf("|*| Environment variable: MCPORT=%d\n", port)
			DefaultPort = port
		}
	}
	// MCHEF
	if echef := os.Getenv("MCHEF"); len(echef) > 0 {
		log.Printf("|*| Environment variable: MCHEF=%s\n", echef)
		DefaultChef = echef
	}
}
