package main

// ====================
//  IMPORTS
// ====================

import (
	masterchef "github.com/cosasdepuma/masterchef/pkg"
)

// ====================
//  PRIVATE METHODS
// ====================

func main() {
	mc := masterchef.New()
	defer mc.Close()
	mc.Start()
}
