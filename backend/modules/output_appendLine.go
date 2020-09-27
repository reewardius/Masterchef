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

type moduleOutputAppendLine Module

var OutputAppendLine = &moduleOutputAppendLine{}

// ====================
//  MODULE METHODS
// ====================

// ----- Normal cooker -----

func (moduleOutputAppendLine) Cook(input string, cals calories) (string, error) {
	text, ok := cals["Text"]
	if !ok {
		return "", fmt.Errorf("Text not specified")
	}
	return fmt.Sprintf("%s\n%s", input, text), nil
}

// ----- Incognito cooker -----

func (m moduleOutputAppendLine) CookShh(input string) (string, error) {
	return "", fmt.Errorf("Icognito mode is not available")
}

// ----- HTML representation -----

func (moduleOutputAppendLine) ToHTML() string {
	// TODO
	return ""
}
