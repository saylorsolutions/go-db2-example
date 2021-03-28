package main

import "testing"

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
