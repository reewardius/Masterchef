package pkg

// ====================
//  IMPORTS
// ====================

import (
	"context"
	"net/http"
	"text/template"

	"github.com/cosasdepuma/masterchef/pkg/internal"
	"github.com/cosasdepuma/masterchef/pkg/public"
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
		Server  *internal.Server
		// Foody
		Kitchen Kitchen
	}
	Kitchen struct {
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
	// -- Handler
	var handler http.Handler
	src, err := template.New("index").Parse(public.Source)
	ok = err == nil
	if err == nil {
		handler = internal.NewRouter(src, channels.GreenLight)
	}
	// -- Server
	srv := internal.NewServer(argv.Host, argv.Port, handler)
	ok = ok && srv != nil
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
