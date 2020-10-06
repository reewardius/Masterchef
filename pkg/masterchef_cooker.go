package pkg

// ====================
//  IMPORTS
// ====================

import (
	"github.com/cosasdepuma/masterchef/pkg/core"
)

// ====================
//  PRIVATE CONSTRUCTOR
// ====================

func newCooker(host string, port int, chef string) (*core.CookerServer, bool) {
	srv := core.NewCookerServer(host, port, chef)
	return srv, srv != nil
}
