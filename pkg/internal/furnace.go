package internal

import (
	"log"
	"sync"

	"github.com/cosasdepuma/masterchef/pkg/modules"
	"github.com/cosasdepuma/masterchef/pkg/utils"
)

func cookInFurnace(data []byte, opts map[string]interface{}) []string {
	// Get how to cook the dish
	dish := modules.NewDish(data)
	// Get the first ingredient (input)
	input := utils.SplitContentLines(dish.Input)
	// Cook step by step
	for i := range dish.Recipes {
		// Prepare current step
		errs := []string{}
		result := []string{}
		module := dish.Recipes[i]
		dish.Recipes[i].Input = input
		// Concurrency
		wg := sync.WaitGroup{}
		wg.Add(len(input))
		lock := sync.Mutex{}
		// Split targets
		for _, target := range input {
			go func(module modules.Recipe, target string) {
				defer wg.Done()
				var err error
				var output []string
				recipe, ok := modules.Recipes[module.Module]
				if !ok {
					log.Printf("|!| Cannot fin module \"%s\"\n", module.Module)
					return
				}
				if module.Incognito {
					output, err = recipe.CookShh(target, opts)
				} else {
					output, err = recipe.Cook(target, module.Arguments, opts)
				}
				lock.Lock()
				// Check errors
				if err != nil {
					log.Printf("|*| Error in module %s: %s", module.Module, err.Error())
					errs = append(errs, err.Error())
				}
				// Check output
				if len(output) > 0 {
					result = append(result, output...)
				}
				lock.Unlock()
			}(module, target)
		}
		wg.Wait()
		// Prepare next step
		if dish.Recipes[i].Single {
			dish.Recipes[i].Output = append(input, result...)
		} else {
			dish.Recipes[i].Output = result
		}
		if len(input) == 0 {
			input = errs
			dish.Recipes[i].Output = errs
			break
		}
		input = dish.Recipes[i].Output
	}
	return input
}
