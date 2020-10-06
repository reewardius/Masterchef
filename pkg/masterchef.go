package pkg

// ====================
//  IMPORTS
// ====================

import (
	"context"
	"net"
	"time"

	"github.com/cosasdepuma/masterchef/pkg/internal"
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
		// Internal
		Argv    *internal.Arguments
		Channel *internal.Channels
		Server  internal.Server
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
	internal.GetEnvironmentConfig()
	// -- Arguments
	argv := internal.NewArguments()
	// -- Context
	ctx, cancel := context.WithCancel(context.Background())
	// -- Channels
	channels := internal.NewChannels()
	// -- Server
	var ok bool
	var srv internal.Server
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
		// Internal
		Argv:    argv,
		Channel: channels,
		Server:  srv,
		// Foody
		Kitchen: Kitchen{
			Cookers: []string{},
		},
	}
}

// ====================
//  PUBLIC METHODS
// ====================

func (mc *Masterchef) Start() {
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
	for _, cooker := range mc.Kitchen.Cookers {
		conn, err := net.DialTimeout("tcp", cooker, time.Second*10)
		if err == nil {
			conn.Write([]byte("fired!"))
			conn.Close()
		}
	}
	select {
	case <-mc.Ctx.Done():
		mc.Channel.Close()
	}
}
