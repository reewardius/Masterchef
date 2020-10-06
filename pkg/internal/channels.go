package internal

// ====================
//  IMPORTS
// ====================

import (
	"os"
	"os/signal"
)

// ====================
//  TYPES
// ====================

type Channels struct {
	GreenLight chan string
	RedLight   chan os.Signal
}

// ====================
//  CONSTRUCTOR
// ====================

func NewChannels() *Channels {
	// Configuration
	// -- Interrupt trap
	red := make(chan os.Signal, 1)
	signal.Notify(red, os.Interrupt)
	// Channels
	return &Channels{
		GreenLight: make(chan string, 10),
		RedLight:   red,
	}
}

// ====================
//  STRUCTURE METHODS
// ====================

func (ch *Channels) Close() {
	close(ch.GreenLight)
	close(ch.RedLight)
}
