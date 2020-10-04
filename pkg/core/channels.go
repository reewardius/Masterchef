package core

import (
	"os"
	"os/signal"
)

type Channels struct {
	RedLight chan os.Signal
}

func NewChannels() *Channels {
	// Configuration
	// -- Interrupt trap
	red := make(chan os.Signal, 1)
	signal.Notify(red, os.Interrupt)
	// Channels
	return &Channels{
		RedLight: red,
	}
}

func (ch *Channels) Close() {
	close(ch.RedLight)
}
