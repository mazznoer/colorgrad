package colorgrad

import (
	"testing"
)

func TestPreset(t *testing.T) {
	testGrad(t, BrBG(), "#543005", "#003c30")
	testGrad(t, PRGn(), "#40004b", "#00441b")
	testGrad(t, PiYG(), "#8e0152", "#276419")
	testGrad(t, PuOr(), "#7f3b08", "#2d004b")
	testGrad(t, RdBu(), "#67001f", "#053061")
	testGrad(t, RdGy(), "#67001f", "#1a1a1a")
	testGrad(t, RdYlBu(), "#a50026", "#313695")
	testGrad(t, RdYlGn(), "#a50026", "#006837")
	testGrad(t, Spectral(), "#9e0142", "#5e4fa2")

	testGrad(t, Blues(), "#f7fbff", "#08306b")
	testGrad(t, Greens(), "#f7fcf5", "#00441b")
	testGrad(t, Greys(), "#ffffff", "#000000")
	testGrad(t, Oranges(), "#fff5eb", "#7f2704")
	testGrad(t, Purples(), "#fcfbfd", "#3f007d")
	testGrad(t, Reds(), "#fff5f0", "#67000d")

	testGrad(t, Viridis(), "#440154", "#fee825")
	testGrad(t, Inferno(), "#000004", "#fcffa4")
	testGrad(t, Magma(), "#000004", "#fcfdbf")
	testGrad(t, Plasma(), "#0d0887", "#f0f921")

	testGrad(t, BuGn(), "#f7fcfd", "#00441b")
	testGrad(t, BuPu(), "#f7fcfd", "#4d004b")
	testGrad(t, GnBu(), "#f7fcf0", "#084081")
	testGrad(t, OrRd(), "#fff7ec", "#7f0000")
	testGrad(t, PuBuGn(), "#fff7fb", "#014636")
	testGrad(t, PuBu(), "#fff7fb", "#023858")
	testGrad(t, PuRd(), "#f7f4f9", "#67001f")
	testGrad(t, RdPu(), "#fff7f3", "#49006a")
	testGrad(t, YlGnBu(), "#ffffd9", "#081d58")
	testGrad(t, YlGn(), "#ffffe5", "#004529")
	testGrad(t, YlOrBr(), "#ffffe5", "#662506")
	testGrad(t, YlOrRd(), "#ffffcc", "#800026")
}

func testGrad(t *testing.T, grad Gradient, start, end string) {
	if isZeroGradient(grad) {
		t.Errorf("Grad is zeroGradient")
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
