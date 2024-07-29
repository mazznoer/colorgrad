package colorgrad

import (
	"testing"
)

func TestGradientFn(t *testing.T) {
	testFn(t, Cividis())
	testFn(t, Sinebow())
	testFn(t, Turbo())

	testFn(t, CubehelixDefault())
	testFn(t, Cool())
	testFn(t, Warm())
	testFn(t, Rainbow())
}

func TestCyclicalGradient(t *testing.T) {
	var grad Gradient

	grad = Rainbow()
	test(t, grad.At(0).HexString(), grad.At(1).HexString())

	grad = Sinebow()
	test(t, grad.At(0).HexString(), grad.At(1).HexString())
}

func testFn(t *testing.T, grad Gradient) {
	testTrue(t, !isZeroGradient(grad))
	test(t, len(grad.Colors(9)), 9)
}
