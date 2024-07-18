package colorgrad

import (
	"image/color"
	"testing"
)

func TestDomain(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("yellow", "blue", "lime").
		Domain(0, 100).
		Build()

	testStr(t, grad.At(0).HexString(), "#ffff00")
	testStr(t, grad.At(50).HexString(), "#0000ff")
	testStr(t, grad.At(100).HexString(), "#00ff00")
	// outside domain
	testStr(t, grad.At(-10).HexString(), "#ffff00")
	testStr(t, grad.At(150).HexString(), "#00ff00")

	grad, _ = NewGradient().
		HtmlColors("yellow", "blue", "lime").
		Domain(-1, 1).
		Build()

	testStr(t, grad.At(-1).HexString(), "#ffff00")
	testStr(t, grad.At(0).HexString(), "#0000ff")
	testStr(t, grad.At(1).HexString(), "#00ff00")
	// outside domain
	testStr(t, grad.At(-1.5).HexString(), "#ffff00")
	testStr(t, grad.At(1.5).HexString(), "#00ff00")

	grad, _ = NewGradient().
		HtmlColors("#00ff00", "#ff0000", "#ffff00").
		Domain(0, 0.75, 1).
		Build()

	testStr(t, grad.At(0).HexString(), "#00ff00")
	testStr(t, grad.At(0.75).HexString(), "#ff0000")
	testStr(t, grad.At(1).HexString(), "#ffff00")
	// outside domain
	testStr(t, grad.At(-1).HexString(), "#00ff00")
	testStr(t, grad.At(1.5).HexString(), "#ffff00")

	grad, _ = NewGradient().
		HtmlColors("#00ff00", "#ff0000", "#0000ff", "#ffff00").
		Domain(15, 25, 29, 63).
		Build()

	testStr(t, grad.At(15).HexString(), "#00ff00")
	testStr(t, grad.At(25).HexString(), "#ff0000")
	testStr(t, grad.At(29).HexString(), "#0000ff")
	testStr(t, grad.At(63).HexString(), "#ffff00")
	// outside domain
	testStr(t, grad.At(10).HexString(), "#00ff00")
	testStr(t, grad.At(67).HexString(), "#ffff00")
}

func TestGetColors(t *testing.T) {
	grad, _ := NewGradient().Build()
	colorsA := grad.ColorfulColors(5) // []colorful.Color
	colorsB := grad.Colors(5)         // []color.Color

	for i, c2 := range colorsB {
		var c1 color.Color = colorsA[i]
		if c1 != c2 {
			t.Errorf("%v != %v", c1, c2)
		}
	}

	colors0 := grad.ColorfulColors(0)
	if len(colors0) != 0 {
		t.Errorf("Error.")
	}
	//colors1 := grad.ColorfulColors(1)
	//testStr(t, colors1[0].HexString(), "#000000")

	colors2 := grad.ColorfulColors(2)
	testStr(t, colors2[0].HexString(), "#000000")
	testStr(t, colors2[1].HexString(), "#ffffff")

	colors3 := grad.ColorfulColors(3)
	testStr(t, colors3[0].HexString(), "#000000")
	testStr(t, colors3[1].HexString(), "#808080")
	testStr(t, colors3[2].HexString(), "#ffffff")

	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Domain(-1, 1).
		Build()

	colors5 := grad.ColorfulColors(5)
	testStr(t, colors5[0].HexString(), "#ff0000")
	testStr(t, colors5[1].HexString(), "#808000")
	testStr(t, colors5[2].HexString(), "#00ff00")
	testStr(t, colors5[3].HexString(), "#008080")
	testStr(t, colors5[4].HexString(), "#0000ff")
}

func TestSpreadRepeat(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("#000", "#fff").
		Build()

	testStr(t, grad.RepeatAt(-2.0).HexString(), "#000000")
	testStr(t, grad.RepeatAt(-1.9).HexString(), "#1a1a1a")
	testStr(t, grad.RepeatAt(-1.5).HexString(), "#808080")
	testStr(t, grad.RepeatAt(-1.1).HexString(), "#e5e5e5")

	testStr(t, grad.RepeatAt(-1.0).HexString(), "#000000")
	testStr(t, grad.RepeatAt(-0.9).HexString(), "#191919")
	testStr(t, grad.RepeatAt(-0.5).HexString(), "#808080")
	testStr(t, grad.RepeatAt(-0.1).HexString(), "#e6e6e6")

	testStr(t, grad.RepeatAt(0.0).HexString(), "#000000")
	testStr(t, grad.RepeatAt(0.1).HexString(), "#1a1a1a")
	testStr(t, grad.RepeatAt(0.5).HexString(), "#808080")
	testStr(t, grad.RepeatAt(0.9).HexString(), "#e5e5e5")

	testStr(t, grad.RepeatAt(1.0).HexString(), "#000000")
	testStr(t, grad.RepeatAt(1.1).HexString(), "#1a1a1a")
	testStr(t, grad.RepeatAt(1.5).HexString(), "#808080")
	testStr(t, grad.RepeatAt(1.9).HexString(), "#e5e5e5")

	testStr(t, grad.RepeatAt(2.0).HexString(), "#000000")
	testStr(t, grad.RepeatAt(2.1).HexString(), "#1a1a1a")
	testStr(t, grad.RepeatAt(2.5).HexString(), "#808080")
	testStr(t, grad.RepeatAt(2.9).HexString(), "#e5e5e5")
}

func TestSpreadReflect(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("#000", "#fff").
		Build()

	testStr(t, grad.ReflectAt(-2.0).HexString(), "#000000")
	testStr(t, grad.ReflectAt(-1.9).HexString(), "#1a1a1a")
	testStr(t, grad.ReflectAt(-1.5).HexString(), "#808080")
	testStr(t, grad.ReflectAt(-1.1).HexString(), "#e5e5e5")

	testStr(t, grad.ReflectAt(-1.0).HexString(), "#ffffff")
	testStr(t, grad.ReflectAt(-0.9).HexString(), "#e5e5e5")
	testStr(t, grad.ReflectAt(-0.5).HexString(), "#808080")
	testStr(t, grad.ReflectAt(-0.1).HexString(), "#1a1a1a")

	testStr(t, grad.ReflectAt(0.0).HexString(), "#000000")
	testStr(t, grad.ReflectAt(0.1).HexString(), "#1a1a1a")
	testStr(t, grad.ReflectAt(0.5).HexString(), "#808080")
	testStr(t, grad.ReflectAt(0.9).HexString(), "#e5e5e5")

	testStr(t, grad.ReflectAt(1.0).HexString(), "#ffffff")
	testStr(t, grad.ReflectAt(1.1).HexString(), "#e5e5e5")
	testStr(t, grad.ReflectAt(1.5).HexString(), "#808080")
	testStr(t, grad.ReflectAt(1.9).HexString(), "#1a1a1a")

	testStr(t, grad.ReflectAt(2.0).HexString(), "#000000")
	testStr(t, grad.ReflectAt(2.1).HexString(), "#1a1a1a")
	testStr(t, grad.ReflectAt(2.5).HexString(), "#808080")
	testStr(t, grad.ReflectAt(2.9).HexString(), "#e5e5e5")
}

func testStr(t *testing.T, a, b string) {
	if a != b {
		t.Helper()
		t.Errorf("Left: %s, right: %s", a, b)
	}
}

func isZeroGradient(grad Gradient) bool {
	dmin, dmax := grad.Domain()
	if dmin != 0 || dmax != 1 {
		return false
	}
	colors := grad.ColorfulColors(13)
	black := Color{R: 0, G: 0, B: 0, A: 0}
	for _, col := range colors {
		if col != black {
			return false
		}
	}
	return true
}
