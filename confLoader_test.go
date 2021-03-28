package main

import (
	"testing"

	testify "github.com/stretchr/testify/require"
)

func TestConnParams(t *testing.T) {
	tests := map[string]struct {
		input    map[string]string
		expected string
	}{
		"Single element": {
			input: map[string]string{
				"key": "value",
			},
			expected: "key=value",
		},
		"Two elements": {
			input: map[string]string{
				"first":  "value",
				"second": "value",
			},
			expected: "first=value;second=value",
		},
		"Nil map": {
			input:    nil,
			expected: "",
		},
		"Empty map": {
			input:    map[string]string{},
			expected: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc := tc
			output := toConnString(tc.input, ";")
			if want, got := tc.expected, output; want != got {
				t.Fatalf("Expected %s\nGot : %s\nWant: %s", want, got, want)
			}
		})
	}
}

func TestFilterInclude(t *testing.T) {
	tests := map[string]struct {
		input        map[string]string
		include      []string
		expectedKeys []string
	}{
		"Empty value map": {
			input:        map[string]string{},
			include:      []string{},
			expectedKeys: []string{},
		},
		"Extra keys": {
			input: map[string]string{
				"a": "value",
				"b": "value",
				"c": "value",
			},
			include:      []string{"a", "b", "c", "d"},
			expectedKeys: []string{"a", "b", "c"},
		},
		"Subset": {
			input: map[string]string{
				"a": "value",
				"b": "value",
				"c": "value",
			},
			include:      []string{"a", "c"},
			expectedKeys: []string{"a", "c"},
		},
		"Include nothing": {
			input: map[string]string{
				"a": "value",
				"b": "value",
				"c": "value",
			},
			include:      []string{},
			expectedKeys: []string{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc := tc
			assert := testify.New(t)

			filtered := filterMapInclude(tc.input, tc.include)
			assert.Len(filtered, len(tc.expectedKeys))
			for _, k := range tc.expectedKeys {
				v, ok := filtered[k]
				assert.True(ok, "Missing key %s", v)
			}
		})
	}
}

func TestFilterExclude(t *testing.T) {
	tests := map[string]struct {
		input        map[string]string
		exclude      []string
		expectedKeys []string
	}{
		"Empty value map": {
			input:        map[string]string{},
			exclude:      []string{},
			expectedKeys: []string{},
		},
		"Extra keys": {
			input: map[string]string{
				"a": "value",
				"b": "value",
				"c": "value",
			},
			exclude:      []string{"a", "b", "c", "d"},
			expectedKeys: []string{},
		},
		"Subset": {
			input: map[string]string{
				"a": "value",
				"b": "value",
				"c": "value",
			},
			exclude:      []string{"a", "c"},
			expectedKeys: []string{"b"},
		},
		"Exclude nothing": {
			input: map[string]string{
				"a": "value",
				"b": "value",
				"c": "value",
			},
			exclude:      []string{},
			expectedKeys: []string{"a", "b", "c"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tc := tc
			assert := testify.New(t)

			filtered := filterMapExclude(tc.input, tc.exclude)
			assert.Len(filtered, len(tc.expectedKeys))
			for _, k := range tc.expectedKeys {
				_, ok := filtered[k]
				assert.True(ok, "Missing expected key %s", k)
			}
		})
	}
}
