package internal

// ====================
//  IMPORTS
// ====================
import "flag"

// ====================
//  TYPES
// ====================

type Arguments struct {
	Host    string
	Port    int
	Threads int
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
	// -- Other
	flag.IntVar(&argv.Threads, "t", DefaultThreads, "")
	flag.IntVar(&argv.Threads, "threads", DefaultThreads, "")
	// Get Values
	flag.Parse()
	return argv
}
