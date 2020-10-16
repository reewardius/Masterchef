package modules

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"

	"github.com/cosasdepuma/masterchef/pkg/utils"
)

// ====================
//  IMPORTS
// ====================

// ====================
//  MODULE DEFINITIONS
// ====================

// ----- Module definition -----

type recipeEnumerationSubdomains Recipe

var EnumerationSubdomains = recipeEnumerationSubdomains{}

// ====================
//  MODULE METHODS
// ====================

// ----- Normal Recipe -----

func (recipeEnumerationSubdomains) Cook(input string, cals calories, opts config) ([]string, error) {
	// Wordlist
	wordlist, ok := cals["Wordlist"]
	if !ok {
		return nil, fmt.Errorf("Wordlist not specified")
	}
	// Threads
	threads := opts["threads"].(int)
	// Read the wordlist
	ext, err := utils.ReadFile(wordlist)
	if err != nil {
		return nil, err
	}
	// Results
	subdomains := []string{}
	// Concurrency
	wg := sync.WaitGroup{}
	wg.Add(len(ext))
	wlock := sync.Mutex{}
	tlock := make(chan struct{}, threads)
	// Execution
	for _, e := range ext {
		tlock <- struct{}{}
		go func(e string) {
			defer wg.Done()
			subdomain := fmt.Sprintf("%s.%s", e, input)
			_, err := net.LookupHost(subdomain)
			if err == nil {
				wlock.Lock()
				subdomains = append(subdomains, subdomain)
				wlock.Unlock()
			}
			<-tlock
		}(e)
	}
	wg.Wait()

	return utils.Unique(subdomains), nil
}

// ----- Incognito Recipe -----

func (recipeEnumerationSubdomains) CookShh(input string, _ config) ([]string, error) {
	// APIs
	apis := []func(string) ([]string, error){
		apiThreatCrowd,
	}
	// Results definition
	results := []string{}
	errs := fmt.Errorf("")
	// Concurrency
	wg := sync.WaitGroup{}
	wg.Add(len(apis))
	lock := sync.Mutex{}
	// Concurrent execution
	for _, api := range apis {
		go func(api func(string) ([]string, error)) {
			defer wg.Done()
			defer lock.Unlock()
			result, err := api(input)
			if err != nil {
				lock.Lock()
				errs = fmt.Errorf("%s\n%s", errs, err)
				return
			}
			lock.Lock()
			results = append(results, result...)
		}(api)
	}
	wg.Wait()
	// Check errors
	if len(results) == 0 {
		return nil, errs
	}
	// Parse data
	return utils.Unique(results), nil
}

// ====================
//  INCOGNITO METHODS
// ====================

func apiThreatCrowd(domain string) ([]string, error) {
	// Data structure
	results := struct {
		Subdomains []string `json:"subdomains"`
	}{}
	addr := fmt.Sprintf("https://www.threatcrowd.org/searchApi/v2/domain/report/?domain=%s",
		domain)
	// Get the data
	body, err := utils.GETBody(addr)
	if err != nil {
		return nil, err
	}
	// Parse the JSON
	err = json.Unmarshal(body, &results)
	if err != nil {
		return nil, fmt.Errorf("ThreadCrowd error: %s", err)
	}
	return results.Subdomains, err
}
