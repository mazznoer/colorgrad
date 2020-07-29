package colorgrad

import (
	"testing"
)

func TestFn(t *testing.T) {
	testFn(t, Cividis())
	testFn(t, Sinebow())
	testFn(t, Turbo())

	testFn(t, CubehelixDefault())
	testFn(t, Cool())
	testFn(t, Warm())
	testFn(t, Rainbow())
}

func testFn(t *testing.T, grad Gradient) {
	n := len(grad.Colors(9))
	if n != 9 {
		t.Errorf("Expected 9, got %v", n)
	}
}
