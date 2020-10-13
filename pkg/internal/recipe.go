package internal

import (
	"encoding/json"
)

// ====================
//  IMPORTS
// ====================

// ====================
//  TYPES
// ====================

type Dish struct {
	Input   string   `json:"input"`
	Recipes []Recipe `json:"recipe"`
}

type Recipe struct {
	Module    string            `json:"name"`
	Incognito bool              `json:"incognito"`
	Arguments map[string]string `json:"calories"`
	Input     string
	Output    []string
	Score     int
}

func newDish(raw []byte) Dish {
	var dish Dish
	json.Unmarshal(raw, &dish)
	return dish
}
