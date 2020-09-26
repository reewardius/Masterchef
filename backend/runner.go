package backend

import (
	"encoding/json"
	"fmt"
	"log"
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
	err := json.Unmarshal(raw, &cooker)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	fmt.Println(cooker.Input)
	return cooker.Input
}
