package colorgrad

import (
	"math"
	"testing"
)

func Test_SharpGradient(t *testing.T) {
	var grad, gradBase Gradient

	gradBase, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Build()

	// Sharp(0)
	grad = gradBase.Sharp(0, 0)
	test(t, grad.At(0.0).HexString(), "#ff0000")
	test(t, grad.At(0.5).HexString(), "#ff0000")
	test(t, grad.At(1.0).HexString(), "#ff0000")

	// Sharp(1)
	grad = gradBase.Sharp(1, 0)
	test(t, grad.At(0.0).HexString(), "#ff0000")
	test(t, grad.At(0.5).HexString(), "#ff0000")
	test(t, grad.At(1.0).HexString(), "#ff0000")

	// Sharp(3)
	grad = gradBase.Sharp(3, 0)
	test(t, grad.At(0.0).HexString(), "#ff0000")
	test(t, grad.At(0.2).HexString(), "#ff0000")

	test(t, grad.At(0.4).HexString(), "#00ff00")
	test(t, grad.At(0.5).HexString(), "#00ff00")
	test(t, grad.At(0.6).HexString(), "#00ff00")

	test(t, grad.At(0.9).HexString(), "#0000ff")
	test(t, grad.At(1.0).HexString(), "#0000ff")

	test(t, grad.At(-0.1).HexString(), "#ff0000")
	test(t, grad.At(1.1).HexString(), "#0000ff")
	test(t, grad.At(math.NaN()).HexString(), "#00000000")

	// Sharp(2)
	gradBase, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Domain(-1, 1).
		Build()

	grad = gradBase.Sharp(2, 0)
	test(t, grad.At(-1.0).HexString(), "#ff0000")
	test(t, grad.At(-0.5).HexString(), "#ff0000")
	test(t, grad.At(-0.1).HexString(), "#ff0000")

	test(t, grad.At(0.1).HexString(), "#0000ff")
	test(t, grad.At(0.5).HexString(), "#0000ff")
	test(t, grad.At(1.0).HexString(), "#0000ff")

	// Smoothness
	gradBase, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Build()

	grad = gradBase.Sharp(0, 0.5)
	test(t, grad.At(0.0).HexString(), "#ff0000")
	test(t, grad.At(0.5).HexString(), "#ff0000")
	test(t, grad.At(1.0).HexString(), "#ff0000")

	grad = gradBase.Sharp(1, 0.5)
	test(t, grad.At(0.0).HexString(), "#ff0000")
	test(t, grad.At(0.5).HexString(), "#ff0000")
	test(t, grad.At(1.0).HexString(), "#ff0000")

	grad = gradBase.Sharp(3, 0.1)

	test(t, grad.At(0).HexString(), "#ff0000")
	test(t, grad.At(0.1).HexString(), "#ff0000")

	test(t, grad.At(1.0/3).HexString(), "#808000")

	test(t, grad.At(0.45).HexString(), "#00ff00")
	test(t, grad.At(0.55).HexString(), "#00ff00")

	test(t, grad.At(1.0/3*2).HexString(), "#008080")

	test(t, grad.At(0.9).HexString(), "#0000ff")
	test(t, grad.At(1).HexString(), "#0000ff")

	test(t, grad.At(-0.01).HexString(), "#ff0000")
	test(t, grad.At(1.01).HexString(), "#0000ff")
	test(t, grad.At(math.NaN()).HexString(), "#00000000")
}
