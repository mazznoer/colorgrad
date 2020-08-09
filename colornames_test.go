package colorgrad

import (
	"testing"

	"github.com/lucasb-eyer/go-colorful"
)

func TestColornames(t *testing.T) {
	for name, want := range testCases {
		got, ok := colornames[name]
		if !ok {
			t.Errorf("Did not find %s", name)
			continue
		}
		if got != want {
			t.Errorf("%s:\ngot  %v\nwant %v", name, got, want)
		}
	}
}

var testCases = map[string]colorful.Color{
	"aqua":    {R: 0, G: 1, B: 1},
	"black":   {R: 0, G: 0, B: 0},
	"blue":    {R: 0, G: 0, B: 1},
	"cyan":    {R: 0, G: 1, B: 1},
	"fuchsia": {R: 1, G: 0, B: 1},
	"lime":    {R: 0, G: 1, B: 0},
	"magenta": {R: 1, G: 0, B: 1},
	"red":     {R: 1, G: 0, B: 0},
	"white":   {R: 1, G: 1, B: 1},
	"yellow":  {R: 1, G: 1, B: 0},
}
