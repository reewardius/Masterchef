package pkg

// ====================
//  IMPORTS
// ====================

import (
	"log"
	"net/http"
	"text/template"

	"github.com/cosasdepuma/masterchef/pkg/utils"

	"github.com/cosasdepuma/masterchef/pkg/internal"
)

// ====================
//  PRIVATE CONSTRUCTOR
// ====================

func newChef(host string, port int, green chan string) (*internal.ChefServer, bool) {
	// -- Handler
	var handler http.Handler
	src, err := template.New("index").Parse(source)
	ok := err == nil
	if err == nil {
		handler = internal.NewRouter(src, green)
	}
	// -- Server
	srv := internal.NewChefServer(host, port, handler)
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
			if _, ok := utils.ContainsString(mc.Kitchen.Cookers, cooker); !ok {
				mc.Kitchen.Cookers = append(mc.Kitchen.Cookers, cooker)
				log.Printf("|+| New cooker added: %s\n", cooker)
				log.Printf("|+| Total cookers: %d\n", len(mc.Kitchen.Cookers))
			}
		case <-mc.Ctx.Done():
			return
		}
	}
}
