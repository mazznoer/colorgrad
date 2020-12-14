package colorgrad

import (
	"testing"
)

func TestLinspace(t *testing.T) {
	result := linspace(0, 1, 2)
	expected := []float64{0, 1}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("%v != %v", result, expected)
			break
		}
	}

	result = linspace(0, 1, 3)
	expected = []float64{0, 0.5, 1}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("%v != %v", result, expected)
			break
		}
	}

	result = linspace(0, 100, 3)
	expected = []float64{0, 50, 100}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("%v != %v", result, expected)
			break
		}
	}
}
