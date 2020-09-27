package backend

// ====================
//  IMPORTS
// ====================

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cosasdepuma/masterchef/backend/modules"
)

// ====================
//  STRUCTS
// ====================

type recipe struct {
	Input  string           `json:"input"`
	Recipe []modules.Module `json:"recipe"`
}

// ====================
//  PUBLIC METHODS
// ====================

// Runner TODO
func Runner(raw []byte) (string, error) {
	cooker := recipe{}
	// Unmarshal recipe
	err := json.Unmarshal(raw, &cooker)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	// Initial values
	result := cooker.Input
	for _, module := range cooker.Recipe {
		ref, ok := modlist[module.Name]
		if !ok {
			return "", fmt.Errorf("%s is not valid a module name", module.Name)
		}
		data, err := ref.Cook(result, module.Calories)
		if err != nil {
			return "", fmt.Errorf("Error running %s:\n\t%s", module.Name, err)
		}
		result = data
	}
	return result, nil
}
