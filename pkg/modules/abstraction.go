package modules

import "encoding/json"

// ====================
//  TYPES
// ====================

type (

	// ----- ALIASES -----

	calories map[string]string
	config   map[string]interface{}

	// ----- INTERFACES -----

	Scheme interface {
		Cook(string, calories, config) ([]string, error)
		CookShh(string, config) ([]string, error)
	}

	// ----- DISHES -----

	Dish struct {
		Input   string   `json:"input"`
		Recipes []Recipe `json:"recipe"`
	}

	Recipe struct {
		Module    string            `json:"name"`
		Single    bool              `json:"single"`
		Incognito bool              `json:"incognito"`
		Arguments map[string]string `json:"calories"`
		Input     []string
		Output    []string
		Score     int
	}
)

// ====================
//  CONSTRUCTOR
// ====================

func NewDish(raw []byte) *Dish {
	var dish Dish
	json.Unmarshal(raw, &dish)
	return &dish
}
