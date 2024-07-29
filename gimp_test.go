package colorgrad

import (
	"math"
	"strings"
	"testing"
)

func TestParseGgr(t *testing.T) {
	black := Rgb(0, 0, 0, 1)
	red := Rgb(1, 0, 0, 1)
	blue := Rgb(0, 0, 1, 1)

	// Black to white
	ggr := "GIMP Gradient\nName: My Gradient\n1\n0 0.5 1 0 0 0 1 1 1 1 1 0 0 0 0"
	grad, name, err := ParseGgr(strings.NewReader(ggr), black, black)
	test(t, err, nil)
	test(t, name, "My Gradient")
	test(t, grad.At(0).HexString(), "#000000")
	test(t, grad.At(1).HexString(), "#ffffff")
	test(t, grad.At(-0.5).HexString(), "#000000")
	test(t, grad.At(1.5).HexString(), "#ffffff")
	test(t, grad.At(math.NaN()).HexString(), "#000000")

	// Foreground to background
	ggr = "GIMP Gradient\nName: My Gradient\n1\n0 0.5 1 0 0 0 1 1 1 1 1 0 0 1 3"
	grad, name, err = ParseGgr(strings.NewReader(ggr), red, blue)
	test(t, err, nil)
	test(t, grad.At(0).HexString(), "#ff0000")
	test(t, grad.At(1).HexString(), "#0000ff")

	// Blending function: step
	ggr = "GIMP Gradient\nName: My Gradient\n1\n0 0.5 1 1 0 0 1 0 0 1 1 5 0 0 0"
	grad, name, err = ParseGgr(strings.NewReader(ggr), black, black)
	test(t, err, nil)
	test(t, grad.At(0.00).HexString(), "#ff0000")
	test(t, grad.At(0.25).HexString(), "#ff0000")
	test(t, grad.At(0.49).HexString(), "#ff0000")
	test(t, grad.At(0.51).HexString(), "#0000ff")
	test(t, grad.At(0.75).HexString(), "#0000ff")
	test(t, grad.At(1.00).HexString(), "#0000ff")

	// Coloring type: HSV CCW (white to blue)
	ggr = "GIMP Gradient\nName: My Gradient\n1\n0 0.5 1 1 1 1 1 0 0 1 1 0 1 0 0"
	grad, name, err = ParseGgr(strings.NewReader(ggr), black, black)
	test(t, err, nil)
	test(t, grad.At(0.0).HexString(), "#ffffff")
	test(t, grad.At(0.5).HexString(), "#80ff80")
	test(t, grad.At(1.0).HexString(), "#0000ff")

	// Coloring type: HSV CW (white to blue)
	ggr = "GIMP Gradient\nName: My Gradient\n1\n0 0.5 1 1 1 1 1 0 0 1 1 0 2 0 0"
	grad, name, err = ParseGgr(strings.NewReader(ggr), black, black)
	test(t, err, nil)
	test(t, grad.At(0.0).HexString(), "#ffffff")
	test(t, grad.At(0.5).HexString(), "#ff80ff")
	test(t, grad.At(1.0).HexString(), "#0000ff")
}

func TestParseGgrError(t *testing.T) {
	black := Rgb(0, 0, 0, 1)
	data := []string{
		"",
		"GIMP Pallete",
		"GIMP Gradient\nxx",
		"GIMP Gradient\nName: Gradient\nx",
		"GIMP Gradient\nName: Gradient\n1\n0 0 0",
	}
	for _, s := range data {
		_, _, err := ParseGgr(strings.NewReader(s), black, black)
		testTrue(t, err != nil)
	}
}
