package backend

// ====================
//  IMPORTS
// ====================

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/cosasdepuma/masterchef/backend/modules"
	"github.com/cosasdepuma/masterchef/backend/utils"
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

// FIXME: Create configuration
const THREADS = 100

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
	input := cooker.Input
	for _, module := range cooker.Recipe {
		// Module
		ref, ok := modlist[module.Name]
		if !ok {
			return "", fmt.Errorf("%s is not valid a module name", module.Name)
		}
		// Inputs
		targets := []string{input}
		if !module.Single {
			targets = strings.Split(strings.TrimSpace(input), "\n")
		}
		// Outputs
		results := ""
		errs := fmt.Errorf("")
		// Concurrency
		wg := sync.WaitGroup{}
		wg.Add(len(targets))
		mlock := sync.Mutex{}
		tlock := make(chan struct{}, THREADS)
		defer close(tlock)
		// Execution
		for _, target := range targets {
			tlock <- struct{}{} // locks a thread
			// Concurrent execution
			go func(target string) {
				defer wg.Done()
				defer mlock.Unlock()
				data, err := "", fmt.Errorf("")
				if module.Incognito {
					data, err = ref.CookShh(target)
				} else {
					data, err = ref.Cook(target, module.Calories)
				}
				<-tlock // releases the thread
				if err != nil {
					mlock.Lock()
					errs = fmt.Errorf("%s\n\t%s", errs, err)
					return
				}
				mlock.Lock()
				results = fmt.Sprintf("%s%s\n", results, data)
			}(target)
		}
		wg.Wait()
		results = strings.TrimSpace(results)
		if len(results) == 0 && len(strings.TrimSpace(fmt.Sprintf("%s", errs))) > 0 {
			return "", fmt.Errorf("Error running %s:\n\t%s", module.Name, errs)
		}
		r := strings.Split(results, "\n")
		r = utils.Unique(r)
		input = utils.ToString(r)
	}
	return input, nil
}
