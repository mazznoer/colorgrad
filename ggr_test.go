package colorgrad

import (
	"image/color"
	"math"
	"strings"
	"testing"
)

func TestParseGgr(t *testing.T) {
	// Black to white
	ggr := "GIMP Gradient\nName: My Gradient\n1\n0 0.5 1 0 0 0 1 1 1 1 1 0 0 0 0"
	grad, name, _ := ParseGgr(strings.NewReader(ggr), color.Black, color.Black)
	testStr(t, name, "My Gradient")
	testStr(t, grad.At(0).Hex(), "#000000")
	testStr(t, grad.At(1).Hex(), "#ffffff")
	testStr(t, grad.At(-0.5).Hex(), "#000000")
	testStr(t, grad.At(1.5).Hex(), "#ffffff")
	testStr(t, grad.At(math.NaN()).Hex(), "#000000")

	// Foreground to background
	ggr = "GIMP Gradient\nName: My Gradient\n1\n0 0.5 1 0 0 0 1 1 1 1 1 0 0 1 3"
	grad, name, _ = ParseGgr(strings.NewReader(ggr), color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255})
	testStr(t, grad.At(0).Hex(), "#ff0000")
	testStr(t, grad.At(1).Hex(), "#0000ff")

	// Blending function: step
	ggr = "GIMP Gradient\nName: My Gradient\n1\n0 0.5 1 1 0 0 1 0 0 1 1 5 0 0 0"
	grad, name, _ = ParseGgr(strings.NewReader(ggr), color.Black, color.Black)
	testStr(t, grad.At(0.00).Hex(), "#ff0000")
	testStr(t, grad.At(0.25).Hex(), "#ff0000")
	testStr(t, grad.At(0.49).Hex(), "#ff0000")
	testStr(t, grad.At(0.51).Hex(), "#0000ff")
	testStr(t, grad.At(0.75).Hex(), "#0000ff")
	testStr(t, grad.At(1.00).Hex(), "#0000ff")

	// Coloring type: HSV CCW (white to blue)
	ggr = "GIMP Gradient\nName: My Gradient\n1\n0 0.5 1 1 1 1 1 0 0 1 1 0 1 0 0"
	grad, name, _ = ParseGgr(strings.NewReader(ggr), color.Black, color.Black)
	testStr(t, grad.At(0.0).Hex(), "#ffffff")
	testStr(t, grad.At(0.5).Hex(), "#80ff80")
	testStr(t, grad.At(1.0).Hex(), "#0000ff")

	// Coloring type: HSV CW (white to blue)
	ggr = "GIMP Gradient\nName: My Gradient\n1\n0 0.5 1 1 1 1 1 0 0 1 1 0 2 0 0"
	grad, name, _ = ParseGgr(strings.NewReader(ggr), color.Black, color.Black)
	testStr(t, grad.At(0.0).Hex(), "#ffffff")
	testStr(t, grad.At(0.5).Hex(), "#ff80ff")
	testStr(t, grad.At(1.0).Hex(), "#0000ff")
}

func TestParseGgrError(t *testing.T) {
	_, _, err0 := ParseGgr(strings.NewReader(""), color.Black, color.Black)
	if err0 == nil {
		t.Error()
	}

	_, _, err1 := ParseGgr(strings.NewReader("GIMP Pallete"), color.Black, color.Black)
	if err1 == nil {
		t.Error()
	}

	_, _, err2 := ParseGgr(strings.NewReader("GIMP Gradient\nxx"), color.Black, color.Black)
	if err2 == nil {
		t.Error()
	}

	_, _, err3 := ParseGgr(strings.NewReader("GIMP Gradient\nName: Gradient\nx"), color.Black, color.Black)
	if err3 == nil {
		t.Error()
	}

	_, _, err4 := ParseGgr(strings.NewReader("GIMP Gradient\nName: Gradient\n1\n0 0 0"), color.Black, color.Black)
	if err4 == nil {
		t.Error()
	}
}
