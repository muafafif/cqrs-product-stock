package model

import (
	"testing"
)

func TestVerifyColor(t *testing.T) {
	tests := []struct {
		name     string
		actual   Product
		expected string
	}{
		{"should pass if the color is red", Product{Color: "Red"}, "Red"},
		{"should pass if the color is green", Product{Color: "Green"}, "Green"},
		{"should pass if the color is blue", Product{Color: "Blue"}, "Blue"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.actual.VerifyColor()
			if err != nil {
				t.Error(err)
			}
			if test.actual.Color != test.expected {
				t.Error("actual and expected not same values")
			}
		})
	}
}
