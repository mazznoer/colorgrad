//go:build ignore
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
	testStr(t, grad.At(0).Hex(), grad.At(1).Hex())

	grad = Sinebow()
	testStr(t, grad.At(0).Hex(), grad.At(1).Hex())
}

func testFn(t *testing.T, grad Gradient) {
	if isZeroGradient(grad) {
		t.Error("grad is zeroGradient")
	}

	n := len(grad.Colors(9))
	if n != 9 {
		t.Errorf("Expected 9, got %v", n)
	}
}
