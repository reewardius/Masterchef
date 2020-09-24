package backend

import (
	"encoding/json"
)

type schm struct {
	Input  string    `json:"input"`
	Recipe []subschm `json:"recipe"`
}

type subschm struct {
	Name     string       `json:"name"`
	Calories []subsubschm `json:"calories"`
}

type subsubschm struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func Runner(raw []byte) string {
	cooker := schm{}
	json.Unmarshal(raw, &cooker)
	return cooker.Input
}
