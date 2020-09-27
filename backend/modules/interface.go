package modules

// ====================
//  STRUCTS
// ====================

type Module struct {
	Name      string            `json:"name"`
	Incognito bool              `json:"incognito"`
	Single    bool              `json:"single"`
	Calories  map[string]string `json:"calories"`
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
	CookShh(string) (string, error)
	ToHTML() string
}
