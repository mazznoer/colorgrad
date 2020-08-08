package colorgrad

import (
	"image/color"
	"math"
	"testing"
)

func TestBasic1(t *testing.T) {
	// Single color
	grad, err := NewGradient().
		Colors(color.RGBA{0, 255, 0, 255}).
		Build()
	testStr(t, grad.At(0).Hex(), "#00ff00")
	testStr(t, grad.At(1).Hex(), "#00ff00")

	// Domain's length != colors's length
	grad, err = NewGradient().
		HtmlColors("#777", "#fff", "#ccc").
		Domain(0, 1).
		Build()
	if err == nil {
		t.Errorf("It should error")
	}
	if grad != nil {
		t.Errorf("grad should nil")
	}

	// Wrong domain
	grad, err = NewGradient().
		HtmlColors("#777", "#fff", "#ccc", "#222").
		Domain(0, 0.71, 0.70, 1).
		Build()
	if err == nil {
		t.Errorf("It should error")
	}
	if grad != nil {
		t.Errorf("grad should nil")
	}

	// Invalid HTML colors
	grad, err = NewGradient().
		HtmlColors("#777", "bloodred", "#bbb", "#zzz").
		Build()
	if err == nil {
		t.Errorf("It should error")
	}
	if grad != nil {
		t.Errorf("grad should nil")
	}

	// Named colors
	grad, err = NewGradient().
		HtmlColors("tomato", "skyblue", "gold", "springgreen").
		Build()
	colors := grad.ColorfulColors(4)
	testStr(t, colors[0].Hex(), "#ff6347")
	testStr(t, colors[1].Hex(), "#87ceeb")
	testStr(t, colors[2].Hex(), "#ffd700")
	testStr(t, colors[3].Hex(), "#00ff7f")

	// Blend mode
	grad, err = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(HCL).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, err = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(HSV).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, err = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(LAB).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, err = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(LRGB).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, err = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(LUV).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")

	grad, err = NewGradient().
		HtmlColors("#333", "#bbb").
		Mode(RGB).
		Build()
	testStr(t, grad.At(0).Hex(), "#333333")
	testStr(t, grad.At(1).Hex(), "#bbbbbb")
}

func TestBasic2(t *testing.T) {
	grad, _ := NewGradient().Build()
	colors := grad.ColorfulColors(2)

	if len(colors) != 2 {
		t.Errorf("Expected 2, got %v", len(colors))
	}
	testStr(t, colors[0].Hex(), "#000000")
	testStr(t, colors[1].Hex(), "#ffffff")

	testStr(t, grad.At(math.NaN()).Hex(), "#000000")

	grad, _ = NewGradient().
		Colors(
			color.RGBA{255, 0, 0, 255},
			color.RGBA{255, 255, 0, 255},
			color.RGBA{0, 0, 255, 255},
		).
		Build()
	colors = grad.ColorfulColors(3)

	if len(colors) != 3 {
		t.Errorf("Expected 3, got %v", len(colors))
	}
	testStr(t, colors[0].Hex(), "#ff0000")
	testStr(t, colors[1].Hex(), "#ffff00")
	testStr(t, colors[2].Hex(), "#0000ff")

	testStr(t, grad.At(math.NaN()).Hex(), "#ff0000")
}

func TestDomain(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("#00ff00", "#ff0000", "#ffff00").
		Domain(0, 0.75, 1).
		Build()

	testStr(t, grad.At(0).Hex(), "#00ff00")
	testStr(t, grad.At(0.75).Hex(), "#ff0000")
	testStr(t, grad.At(1).Hex(), "#ffff00")

	// outside domain
	testStr(t, grad.At(-1).Hex(), "#00ff00")
	testStr(t, grad.At(1.5).Hex(), "#ffff00")

	grad, _ = NewGradient().
		HtmlColors("#00ff00", "#ff0000", "#0000ff", "#ffff00").
		Domain(15, 25, 29, 63).
		Build()

	testStr(t, grad.At(15).Hex(), "#00ff00")
	testStr(t, grad.At(25).Hex(), "#ff0000")
	testStr(t, grad.At(29).Hex(), "#0000ff")
	testStr(t, grad.At(63).Hex(), "#ffff00")

	// outside domain
	testStr(t, grad.At(10).Hex(), "#00ff00")
	testStr(t, grad.At(67).Hex(), "#ffff00")
}

func TestSharp(t *testing.T) {
	grad, _ := NewGradient().Build()
	grad2 := grad.Sharp(7)
	testStr(t, grad2.At(0).Hex(), "#000000")
	testStr(t, grad2.At(1).Hex(), "#ffffff")

	testStr(t, grad2.At(math.NaN()).Hex(), "#000000")
	testStr(t, grad2.At(-0.01).Hex(), "#000000")
	testStr(t, grad2.At(1.01).Hex(), "#ffffff")

	colors := grad2.ColorfulColors(7)
	if len(colors) != 7 {
		t.Errorf("Expected 7, got %v", len(colors))
	}
	testStr(t, colors[0].Hex(), "#000000")
	testStr(t, colors[6].Hex(), "#ffffff")
}

func TestGetColors(t *testing.T) {
	grad, _ := NewGradient().Build()
	colors1 := grad.ColorfulColors(5) // []colorful.Color
	colors2 := grad.Colors(5)         // []color.Color

	for i, c2 := range colors2 {
		var c1 color.Color = colors1[i]
		if c1 != c2 {
			t.Errorf("%v != %v", c1, c2)
		}
	}
}

func testStr(t *testing.T, result, expected string) {
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
