package modules

// ====================
//  IMPORTS
// ====================

import (
	"fmt"
)

// ====================
//  MODULE DEFINITIONS
// ====================

// ----- Module definition -----

var OutputAppend = &Module{}

// ====================
//  MODULE METHODS
// ====================

// ----- Normal cooker -----

func (Module) Cook(input string, cals calories) (string, error) {
	line, ok := cals["Line"]
	if !ok {
		return "", fmt.Errorf("Line not specified")
	}
	return fmt.Sprintf("%s\n%s", input, line), nil
}

// ----- Incognito cooker -----

func (m Module) CookShh(input string, cals calories) (string, error) {
	return m.Cook(input, cals)
}

// ----- HTML representation -----

func (Module) ToHTML() string {
	return ""
}
