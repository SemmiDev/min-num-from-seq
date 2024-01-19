package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestMinimum(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "M", expected: "21"},
		{input: "MM", expected: "321"},
		{input: "N", expected: "12"},
		{input: "NN", expected: "123"},
		{input: "MNMN", expected: "21435"},
		{input: "NNMMM", expected: "126543"},
		{input: "MMNMMNNM", expected: "321654798"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("minimum(%s)", test.input), func(t *testing.T) {
			result := minimum(test.input)
			result = strings.Replace(result, " ", "", -1)

			if result != test.expected {
				t.Errorf("Hasil: %s, Harapannya: %s", result, test.expected)
			}
		})
	}
}
