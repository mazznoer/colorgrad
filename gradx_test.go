package colorgrad

import (
	"image/color"
	"math"
	"testing"
)

func TestX(t *testing.T) {
	grad, _ := NewGradient().Build()
	grad2 := Classes(grad, 7)
	testStr(t, grad2.At(0).Hex(), "#000000")
	testStr(t, grad2.At(1).Hex(), "#ffffff")

	testStr(t, grad2.At(math.NaN()).Hex(), "#000000")
	testStr(t, grad2.At(-0.01).Hex(), "#000000")
	testStr(t, grad2.At(1.01).Hex(), "#ffffff")

	colors := grad2.Colors(7)
	if len(colors) != 7 {
		t.Errorf("Expected 7, got %v", len(colors))
	}
	testStr(t, colors[0].Hex(), "#000000")
	testStr(t, colors[6].Hex(), "#ffffff")

	colors1 := grad.Colors(5)      // []colorful.Color
	colors2 := IntoColors(colors1) // []color.Color

	for i, c2 := range colors2 {
		var c1 color.Color = colors1[i]
		if c1 != c2 {
			t.Errorf("%v != %v", c1, c2)
		}
	}
}
