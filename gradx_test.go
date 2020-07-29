package colorgrad

import (
	"image/color"
	"testing"
)

func TestX(t *testing.T) {
	grad, _ := NewGradient().Build()
	grad2 := Classes(grad, 7)
	testStr(t, grad2.At(0).Hex(), "#000000")
	testStr(t, grad2.At(1).Hex(), "#ffffff")

	colors1 := grad.Colors(5)
	colors2 := IntoColors(colors1)

	for i, c1 := range colors1 {
		var c2 color.Color
		c2 = colors2[i]
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
