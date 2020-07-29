package colorgrad

import (
	"testing"
)

func TestPreset(t *testing.T) {
	testGrad(t, Inferno(), "#000004", "#fcffa4")
	testGrad(t, Magma(), "#000004", "#fcfdbf")
	testGrad(t, Plasma(), "#0d0887", "#f0f921")
	testGrad(t, Spectral(), "#9e0142", "#5e4fa2")
	testGrad(t, Viridis(), "#440154", "#fee825")

	testGrad(t, Blues(), "#f7fbff", "#08306b")
	testGrad(t, Greens(), "#f7fcf5", "#00441b")
	testGrad(t, Greys(), "#ffffff", "#000000")
	testGrad(t, Oranges(), "#fff5eb", "#7f2704")
	testGrad(t, Purples(), "#fcfbfd", "#3f007d")
	testGrad(t, Reds(), "#fff5f0", "#67000d")
}

func testGrad(t *testing.T, grad Gradient, start, end string) {
	if grad == nil {
		t.Errorf("Grad is nil")
	}

	a := grad.At(0).Hex()
	if a != start {
		t.Errorf("Expected %v, got %v", start, a)
	}

	b := grad.At(1).Hex()
	if b != end {
		t.Errorf("Expected %v, got %v", end, b)
	}
}
