package utils

// ====================
//  IMPORTS
// ====================

// ====================
//  PUBLIC METHODS
// ====================

// Unique removes duplicates from a string slice
func Unique(data []string) []string {
	unique := []string{}
	duplicates := make(map[string]int)
	// Iterate over all the data
	for _, d := range data {
		duplicates[d]++
		if duplicates[d] == 1 {
			unique = append(unique, d)
		}
	}
	return unique
}
