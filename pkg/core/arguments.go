package core

// ====================
//  IMPORTS
// ====================
import "flag"

// ====================
//  TYPES
// ====================

type Arguments struct {
	Host string
	Port int
	Chef string
}

// ====================
//  CONSTRUCTOR
// ====================

func NewArguments() *Arguments {
	argv := &Arguments{}
	// Define flags
	// -- Host
	flag.StringVar(&argv.Host, "h", DefaultHost, "")
	flag.StringVar(&argv.Host, "host", DefaultHost, "")
	// -- Port
	flag.IntVar(&argv.Port, "p", DefaultPort, "")
	flag.IntVar(&argv.Port, "port", DefaultPort, "")
	// -- Chef
	flag.StringVar(&argv.Chef, "chef", DefaultChef, "")
	// Get Values
	flag.Parse()
	return argv
}
