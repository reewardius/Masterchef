package pkg

// ====================
//  IMPORTS
// ====================

import (
	"context"

	"github.com/cosasdepuma/masterchef/pkg/core"
)

// ====================
//  STRUCTURES
// ====================

type (
	Masterchef struct {
		// Comunication
		OK        bool
		Ctx       context.Context
		RedButton context.CancelFunc
		// Core
		Argv    *core.Arguments
		Channel *core.Channels
		Server  core.Server
		// Foody
		Kitchen Kitchen
	}
	Kitchen struct {
		Cookers  []string
		Sessions []Session
	}
	Session struct {
		Uuid    string
		Recipes []Result
	}
	Result struct {
		ID        string
		Module    string
		Arguments map[string]string
		Input     string
		Output    []string
		Score     int
	}
)

// ====================
//  PUBLIC CONSTRUCTOR
// ====================

func New() *Masterchef {
	// Configuration
	core.GetEnvironmentConfig()
	// -- Arguments
	argv := core.NewArguments()
	// -- Context
	ctx, cancel := context.WithCancel(context.Background())
	// -- Channels
	channels := core.NewChannels()
	// -- Server
	var ok bool
	var srv core.Server
	if len(argv.Chef) == 0 {
		srv, ok = newChef(argv.Host, argv.Port, channels.GreenLight)
	} else {
		srv, ok = newCooker(argv.Host, argv.Port, argv.Chef)
	}
	// Masterchef
	return &Masterchef{
		// Comunication
		OK:        ok,
		Ctx:       ctx,
		RedButton: cancel,
		// Core
		Argv:    argv,
		Channel: channels,
		Server:  srv,
		Kitchen: Kitchen{
			Cookers: []string{},
		},
	}
}

// ====================
//  PUBLIC METHODS
// ====================

func (mc Masterchef) Start() {
	// Signals
	go func() {
		select {
		case <-mc.Channel.RedLight:
			mc.RedButton()
		}
	}()
	// Cookers
	if len(mc.Argv.Chef) == 0 {
		go mc.chefOrchestration()
	}
	// Server
	mc.Server.Listen(mc.Ctx, mc.RedButton)
}

func (mc *Masterchef) Close() {
	if mc.Ctx.Err() == nil {
		mc.RedButton()
	}
	select {
	case <-mc.Ctx.Done():
		mc.Channel.Close()
	}
}
