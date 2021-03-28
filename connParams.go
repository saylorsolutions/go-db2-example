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
