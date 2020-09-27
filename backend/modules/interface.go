package modules

// ====================
//  STRUCTS
// ====================

type Module struct {
	Name     string            `json:"name"`
	Calories map[string]string `json:"calories"`
}

// ====================
//  ALIASES
// ====================

type calories map[string]string

// ====================
//  INTERFACES
// ====================

type Scheme interface {
	Cook(string, calories) (string, error)
	CookShh(string, calories) (string, error)
	ToHTML() string
}
