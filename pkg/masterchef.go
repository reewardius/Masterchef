package pkg

import (
	"context"
	"log"
	"net/http"
	"text/template"

	"github.com/cosasdepuma/masterchef/pkg/cluster"
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
		Chef    bool
		Kitchen Kitchen
	}
	Kitchen struct {
		Sessions []Session
		Cookers  []string
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
	isCooker := len(argv.Chef) > 0
	var kitchen Kitchen
	if isCooker {
		kitchen.Cookers = []string{argv.Chef}
	}
	// -- Context
	ctx, cancel := context.WithCancel(context.Background())
	// -- Handler
	var handler http.Handler
	if isCooker {
		handler = core.NewRouterCooker()
	} else {
		src, err := template.New("index").Parse(source)
		OK = OK && err == nil
		if err == nil {
			handler = core.NewRouterChef(src)
		}
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
		Chef:    !isCooker,
		Kitchen: kitchen,
	}
}

func (mc *Masterchef) Start() {
	if !mc.Chef {
		// Meet chef
		chef := mc.Kitchen.Cookers[0]
		log.Printf("[*] Cooking for: %s\n", chef)
		if !cluster.MeetChef(chef) {
			log.Println("[!] Where is the chef? There is no masterchef!")
			mc.OK = false
		}
	}
	if mc.OK {
		mc.Server.ListenAndServe(mc.Channel.RedLight)
	}
}

func (mc *Masterchef) Close() {
	mc.Channel.Close()
	mc.RedButton()
}
