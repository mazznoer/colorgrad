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

	colors1 := grad.Colors(5)      // []colorful.Color
	colors2 := IntoColors(colors1) // []color.Color

	for i, c2 := range colors2 {
		var c1 color.Color = colors1[i]
		if c1 != c2 {
			t.Errorf("%v != %v", c1, c2)
		}
	}
}
