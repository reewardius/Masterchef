package pkg

// ====================
//  IMPORTS
// ====================

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"text/template"
	"time"

	"github.com/cosasdepuma/masterchef/pkg/core"
)

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

// ====================
//  STRUCTURE METHODS
// ====================

func (mc *Masterchef) chefOrchestration() {
	for {
		select {
		case cooker := <-mc.Channel.GreenLight:
			if _, ok := mc.Kitchen.Cookers[cooker]; !ok {
				conn, err := net.Dial("tcp", cooker)
				if err == nil {

					mc.Kitchen.Cookers[cooker] = conn
					log.Printf("|+| New cooker added: %s\n", cooker)
					log.Printf("|+| Total cookers: %d\n", len(mc.Kitchen.Cookers))
				}
			}
		case <-time.After(10 * time.Second):
			for addr, cooker := range mc.Kitchen.Cookers {
				cooker.Write([]byte("a coffee"))
				buff := make([]byte, 4)
				_, err := cooker.Read(buff)
				fmt.Println(string(buff))
				if err != nil || string(buff) != "sure" {
					cooker.Close()
					delete(mc.Kitchen.Cookers, addr)
					log.Println(buff, err.Error())
					log.Printf("|-| A cooker was fired: %s\n", addr)
				}
			}
		case <-mc.Ctx.Done():
			return
		}
	}
}
