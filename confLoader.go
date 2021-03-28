package main

import (
	"fmt"
	"log"
	"strings"
)

// toConnString creates a key/value string separating each pair by sep.
func toConnString(values map[string]string, sep string) string {
	if values == nil {
		log.Println("Nil map passed to 'toConnString'")
		return ""
	}

	var pairs []string
	for k, v := range values {
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(pairs, sep)
}

func filterMapInclude(values map[string]string, include []string) map[string]string {
	if values == nil || len(include) == 0 {
		return map[string]string{}
	}
	includeFilter := map[string]bool{}
	for _, v := range include {
		includeFilter[v] = true
	}
	filtered := map[string]string{}
	for k, v := range values {
		if includeFilter[k] {
			filtered[k] = v
		}
	}
	return filtered
}

func filterMapExclude(values map[string]string, exclude []string) map[string]string {
	if values == nil {
		return map[string]string{}
	}
	if len(exclude) == 0 {
		return values
	}
	excludeFilter := map[string]bool{}
	for _, v := range exclude {
		excludeFilter[v] = true
	}
	filtered := map[string]string{}
	for k, v := range values {
		if excludeFilter[k] {
			continue
		}
		filtered[k] = v
	}
	return filtered
}
