package core

import (
	"os"
	"os/signal"
)

type Channels struct {
	GreenLight chan string
	RedLight   chan os.Signal
}

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

func (ch *Channels) Close() {
	close(ch.GreenLight)
	close(ch.RedLight)
}
