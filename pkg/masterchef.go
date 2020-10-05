package pkg

import (
	"context"
	"net/http"
	"text/template"

	"github.com/cosasdepuma/masterchef/pkg/core"
)

type (
	Masterchef struct {
		// Comunication
		OK        bool
		Ctx       context.Context
		RedButton context.CancelFunc
		// Core
		Argv    *core.Arguments
		Channel *core.Channels
		Server  *core.Server
		// Foody
		Kitchen Kitchen
	}
	Kitchen struct {
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

func New() *Masterchef {
	// Configuration
	OK := true
	core.GetEnvironmentConfig()
	// -- Arguments
	argv := core.NewArguments()
	// -- Cookers
	var kitchen Kitchen
	// -- Context
	ctx, cancel := context.WithCancel(context.Background())
	// -- Handler
	var handler http.Handler
	src, err := template.New("index").Parse(source)
	OK = OK && err == nil
	if err == nil {
		handler = core.NewRouterChef(src)
	}
	// -- Server
	srv := core.NewServer(argv.Host, argv.Port, handler)
	OK = OK && srv != nil
	// Masterchef
	return &Masterchef{
		// Comunication
		OK:        srv != nil,
		Ctx:       ctx,
		RedButton: cancel,
		// Core
		Argv:    argv,
		Channel: core.NewChannels(),
		Server:  srv,
		// Foody
		Kitchen: kitchen,
	}
}

func (mc *Masterchef) Start() {
	mc.Server.ListenAndServe(mc.Channel.RedLight)
}

func (mc *Masterchef) Close() {
	mc.Channel.Close()
	mc.RedButton()
}
