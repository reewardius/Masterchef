package pkg

// ====================
//  IMPORTS
// ====================

import (
	"github.com/cosasdepuma/masterchef/pkg/internal"
)

// ====================
//  PRIVATE CONSTRUCTOR
// ====================

func newCooker(host string, port int, chef string) (*internal.CookerServer, bool) {
	srv := internal.NewCookerServer(host, port, chef)
	return srv, srv != nil
}
