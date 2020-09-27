package backend

// ====================
//  IMPORTS
// ====================

import (
	"github.com/cosasdepuma/masterchef/backend/modules"
)

// ====================
//  GLOBALS
// ====================

var modlist = map[string]modules.Scheme{
	// ----- Output -----
	"Append": modules.OutputAppend,
}
