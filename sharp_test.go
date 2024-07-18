package colorgrad

import (
	"math"
	"testing"
)

func TestSharpGradient(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Build()

	// Sharp(0)
	grad0 := grad.Sharp(0, 0)
	testStr(t, grad0.At(0.0).HexString(), "#ff0000")
	testStr(t, grad0.At(0.5).HexString(), "#ff0000")
	testStr(t, grad0.At(1.0).HexString(), "#ff0000")

	// Sharp(1)
	grad1 := grad.Sharp(1, 0)
	testStr(t, grad1.At(0.0).HexString(), "#ff0000")
	testStr(t, grad1.At(0.5).HexString(), "#ff0000")
	testStr(t, grad1.At(1.0).HexString(), "#ff0000")

	// Sharp(3)
	grad3 := grad.Sharp(3, 0)
	testStr(t, grad3.At(0.0).HexString(), "#ff0000")
	testStr(t, grad3.At(0.2).HexString(), "#ff0000")

	testStr(t, grad3.At(0.4).HexString(), "#00ff00")
	testStr(t, grad3.At(0.5).HexString(), "#00ff00")
	testStr(t, grad3.At(0.6).HexString(), "#00ff00")

	testStr(t, grad3.At(0.9).HexString(), "#0000ff")
	testStr(t, grad3.At(1.0).HexString(), "#0000ff")

	testStr(t, grad3.At(-0.1).HexString(), "#ff0000")
	testStr(t, grad3.At(1.1).HexString(), "#0000ff")
	testStr(t, grad3.At(math.NaN()).HexString(), "#00000000")

	// Sharp(2)
	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Domain(-1, 1).
		Build()
	grad2 := grad.Sharp(2, 0)
	testStr(t, grad2.At(-1.0).HexString(), "#ff0000")
	testStr(t, grad2.At(-0.5).HexString(), "#ff0000")
	testStr(t, grad2.At(-0.1).HexString(), "#ff0000")

	testStr(t, grad2.At(0.1).HexString(), "#0000ff")
	testStr(t, grad2.At(0.5).HexString(), "#0000ff")
	testStr(t, grad2.At(1.0).HexString(), "#0000ff")
}

func TestSharpGradientSmoothness(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Build()

	grad0 := grad.Sharp(0, 0.5)
	testStr(t, grad0.At(0.0).HexString(), "#ff0000")
	testStr(t, grad0.At(0.5).HexString(), "#ff0000")
	testStr(t, grad0.At(1.0).HexString(), "#ff0000")

	grad1 := grad.Sharp(1, 0.5)
	testStr(t, grad1.At(0.0).HexString(), "#ff0000")
	testStr(t, grad1.At(0.5).HexString(), "#ff0000")
	testStr(t, grad1.At(1.0).HexString(), "#ff0000")

	grad = grad.Sharp(3, 0.1)

	testStr(t, grad.At(0).HexString(), "#ff0000")
	testStr(t, grad.At(0.1).HexString(), "#ff0000")

	testStr(t, grad.At(1.0/3).HexString(), "#808000")

	testStr(t, grad.At(0.45).HexString(), "#00ff00")
	testStr(t, grad.At(0.55).HexString(), "#00ff00")

	testStr(t, grad.At(1.0/3*2).HexString(), "#008080")

	testStr(t, grad.At(0.9).HexString(), "#0000ff")
	testStr(t, grad.At(1).HexString(), "#0000ff")

	testStr(t, grad.At(-0.01).HexString(), "#ff0000")
	testStr(t, grad.At(1.01).HexString(), "#0000ff")
	testStr(t, grad.At(math.NaN()).HexString(), "#00000000")
}
