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
	// ----- Enumeration -----
	"Subdomains": modules.EnumerationSubdomains,
	// ----- Output -----
	"Append Line": modules.OutputAppendLine,
}
