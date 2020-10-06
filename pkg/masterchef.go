package pkg

// ====================
//  IMPORTS
// ====================

import (
	"context"
	"log"
	"net"
	"net/http"
	"text/template"
	"time"

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
		Cookers  map[string]net.Conn
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
			Cookers: map[string]net.Conn{},
		},
	}
}

// ====================
//  PRIVATE CONSTRUCTOR
// ====================

func newChef(host string, port int, green chan string) (*core.ChefServer, bool) {
	// -- Handler
	var handler http.Handler
	src, err := template.New("index").Parse(source)
	ok := err == nil
	if err == nil {
		handler = core.NewRouter(src, green)
	}
	// -- Server
	srv := core.NewChefServer(host, port, handler)
	ok = ok && srv != nil
	return srv, ok
}

func newCooker(host string, port int, chef string) (*core.CookerServer, bool) {
	srv := core.NewCookerServer(host, port, chef)
	return srv, srv != nil
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
		go func() {
			for {
				select {
				case <-time.After(10 * time.Second):
					for addr, cooker := range mc.Kitchen.Cookers {
						cooker.Write([]byte("a coffee"))
						buff := make([]byte, 4)
						_, err := cooker.Read(buff)
						if err != nil || string(buff) != "sure" {
							cooker.Close()
							delete(mc.Kitchen.Cookers, addr)
							log.Println(buff, err.Error())
							log.Printf("|-| A cooker was fired: %s\n", addr)
						}
					}
				case cooker := <-mc.Channel.GreenLight:
					if _, ok := mc.Kitchen.Cookers[cooker]; !ok {
						conn, err := net.Dial("tcp", cooker)
						if err == nil {
							err := conn.SetReadDeadline(time.Now().Add(5 * time.Second))
							if err == nil {
								mc.Kitchen.Cookers[cooker] = conn
								log.Printf("|+| New cooker added: %s\n", cooker)
								log.Printf("|+| Total cookers: %d\n", len(mc.Kitchen.Cookers))
							}
						}
					}
				case <-mc.Ctx.Done():
					return
				}
			}
		}()
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
		for _, cooker := range mc.Kitchen.Cookers {
			cooker.Close()
		}
		mc.Channel.Close()
	}
}
