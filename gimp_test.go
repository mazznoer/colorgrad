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
	grad, name, _ := ParseGgr(strings.NewReader(ggr), black, black)
	testStr(t, name, "My Gradient")
	testStr(t, grad.At(0).HexString(), "#000000")
	testStr(t, grad.At(1).HexString(), "#ffffff")
	testStr(t, grad.At(-0.5).HexString(), "#000000")
	testStr(t, grad.At(1.5).HexString(), "#ffffff")
	testStr(t, grad.At(math.NaN()).HexString(), "#000000")

	// Foreground to background
	ggr = "GIMP Gradient\nName: My Gradient\n1\n0 0.5 1 0 0 0 1 1 1 1 1 0 0 1 3"
	grad, name, _ = ParseGgr(strings.NewReader(ggr), red, blue)
	testStr(t, grad.At(0).HexString(), "#ff0000")
	testStr(t, grad.At(1).HexString(), "#0000ff")

	// Blending function: step
	ggr = "GIMP Gradient\nName: My Gradient\n1\n0 0.5 1 1 0 0 1 0 0 1 1 5 0 0 0"
	grad, name, _ = ParseGgr(strings.NewReader(ggr), black, black)
	testStr(t, grad.At(0.00).HexString(), "#ff0000")
	testStr(t, grad.At(0.25).HexString(), "#ff0000")
	testStr(t, grad.At(0.49).HexString(), "#ff0000")
	testStr(t, grad.At(0.51).HexString(), "#0000ff")
	testStr(t, grad.At(0.75).HexString(), "#0000ff")
	testStr(t, grad.At(1.00).HexString(), "#0000ff")

	// Coloring type: HSV CCW (white to blue)
	ggr = "GIMP Gradient\nName: My Gradient\n1\n0 0.5 1 1 1 1 1 0 0 1 1 0 1 0 0"
	grad, name, _ = ParseGgr(strings.NewReader(ggr), black, black)
	testStr(t, grad.At(0.0).HexString(), "#ffffff")
	testStr(t, grad.At(0.5).HexString(), "#80ff80")
	testStr(t, grad.At(1.0).HexString(), "#0000ff")

	// Coloring type: HSV CW (white to blue)
	ggr = "GIMP Gradient\nName: My Gradient\n1\n0 0.5 1 1 1 1 1 0 0 1 1 0 2 0 0"
	grad, name, _ = ParseGgr(strings.NewReader(ggr), black, black)
	testStr(t, grad.At(0.0).HexString(), "#ffffff")
	testStr(t, grad.At(0.5).HexString(), "#ff80ff")
	testStr(t, grad.At(1.0).HexString(), "#0000ff")
}

func TestParseGgrError(t *testing.T) {
	black := Rgb(0, 0, 0, 1)

	_, _, err0 := ParseGgr(strings.NewReader(""), black, black)
	if err0 == nil {
		t.Error()
	}

	_, _, err1 := ParseGgr(strings.NewReader("GIMP Pallete"), black, black)
	if err1 == nil {
		t.Error()
	}

	_, _, err2 := ParseGgr(strings.NewReader("GIMP Gradient\nxx"), black, black)
	if err2 == nil {
		t.Error()
	}

	_, _, err3 := ParseGgr(strings.NewReader("GIMP Gradient\nName: Gradient\nx"), black, black)
	if err3 == nil {
		t.Error()
	}

	_, _, err4 := ParseGgr(strings.NewReader("GIMP Gradient\nName: Gradient\n1\n0 0 0"), black, black)
	if err4 == nil {
		t.Error()
	}
}
