package modules

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"

	"github.com/cosasdepuma/masterchef/backend/utils"
)

// ====================
//  IMPORTS
// ====================

// ====================
//  MODULE DEFINITIONS
// ====================

// ----- Module definition -----

type moduleEnumerationSubdomains Module

var EnumerationSubdomains = &moduleEnumerationSubdomains{}

// ====================
//  MODULE METHODS
// ====================

// ----- Normal cooker -----

func (moduleEnumerationSubdomains) Cook(input string, cals calories) (string, error) {
	// Wordlist
	wordlist, ok := cals["Wordlist"]
	if !ok {
		return "", fmt.Errorf("Wordlist not specified")
	}
	// Read the wordlist
	ext, err := utils.ReadFile(wordlist)
	if err != nil {
		return "", err
	}
	// Results
	subdomains := []string{}
	// Concurrency
	wg := sync.WaitGroup{}
	wg.Add(len(ext))
	wlock := sync.Mutex{}
	tlock := make(chan struct{}, 7) // FIXME Change number of threads
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

	subdomains = utils.Unique(subdomains)
	result := utils.ToString(subdomains)
	return result, nil
}

// ----- Incognito cooker -----

func (moduleEnumerationSubdomains) CookShh(input string) (string, error) {
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
		go func() {
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
		}()
	}
	wg.Wait()
	// Check errors
	if len(results) == 0 {
		return "", errs
	}
	// Parse data
	results = utils.Unique(results)
	return utils.ToString(results), nil
}

// ----- HTML representation -----

func (moduleEnumerationSubdomains) ToHTML() string {
	// TODO
	return ""
}

// ====================
//  INCOGNITO METHODS
// ====================

func apiThreatCrowd(domain string) ([]string, error) {
	// Data structure
	results := struct {
		Subdomains []string `json:"subdomains"`
	}{}
	// Get the data
	body, err := utils.Get("https://www.threatcrowd.org/searchApi/v2/domain/report/?domain=%s",
		domain)
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
