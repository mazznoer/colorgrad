package colorgrad

import (
	"testing"

	"github.com/lucasb-eyer/go-colorful"
)

func TestParseColor(t *testing.T) {
	type colorPair struct {
		in  string
		out colorful.Color
	}

	testData := []colorPair{
		{"red", colorful.Color{R: 1, G: 0, B: 0}},
		{"#f00", colorful.Color{R: 1, G: 0, B: 0}},
		{"#ff0000", colorful.Color{R: 1, G: 0, B: 0}},
		{"rgb(255,0,0)", colorful.Color{R: 1, G: 0, B: 0}},
		{"RGB( 255 , 255 , 0 )", colorful.Color{R: 1, G: 1, B: 0}},
		{"rgba(255,0,0,0.5)", colorful.Color{R: 1, G: 0, B: 0}},
		{"rgb(100%,0%,0%)", colorful.Color{R: 1, G: 0, B: 0}},
		{"rgba(100%,0%,0%,0.5)", colorful.Color{R: 1, G: 0, B: 0}},
		{"hsl(0,100%,50%)", colorful.Color{R: 1, G: 0, B: 0}},
		{"hsla(360,100%,50%,0.5)", colorful.Color{R: 1, G: 0, B: 0}},
	}

	for _, d := range testData {
		c, err := parseColor(d.in)
		if err != nil {
			t.Errorf("Parse error: %s", d.in)
			continue
		}
		if c != d.out {
			t.Errorf("%s -> %v != %v", d.in, c, d.out)
		}
	}
}

func TestLinspace(t *testing.T) {
	result := linspace(0, 1, 2)
	expected := []float64{0, 1}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("%v != %v", result, expected)
		}
	}

	result = linspace(0, 1, 3)
	expected = []float64{0, 0.5, 1}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("%v != %v", result, expected)
		}
	}

	result = linspace(0, 100, 3)
	expected = []float64{0, 50, 100}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("%v != %v", result, expected)
		}
	}
}
