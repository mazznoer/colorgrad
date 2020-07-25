package colorgrad

import (
    "image/color"
    "testing"
)

func TestBasic(t *testing.T) {
    grad, _ := NewGradient().Build()
    colors := grad.Colors(2)

    if len(colors) != 2 {
        t.Errorf("Expected 2, got %v", len(colors))
    }
    testStr(t, colors[0].Hex(), "#000000")
    testStr(t, colors[1].Hex(), "#ffffff")

    grad, _ = NewGradient().
        Colors(
            color.RGBA{255,0,0,255},
            color.RGBA{255,255,0,255},
            color.RGBA{0,0,255,255},
        ).
        Build()

    colors = grad.Colors(3)

    if len(colors) != 3 {
        t.Errorf("%v", len(colors))
    }
    testStr(t, colors[0].Hex(), "#ff0000")
    testStr(t, colors[1].Hex(), "#ffff00")
    testStr(t, colors[2].Hex(), "#0000ff")
}

func TestDomain(t *testing.T) {
    grad, _ := NewGradient().
        HexColors("#000000", "#ff0000", "#ffffff").
        Domain(0, 0.75, 1).
        Build()
    
    testStr(t, grad.At(0).Hex(), "#000000")
    testStr(t, grad.At(0.75).Hex(), "#ff0000")
    testStr(t, grad.At(1).Hex(), "#ffffff")
    
    grad, _ = NewGradient().
        HexColors("#000000", "#ff0000", "#ffffff").
        Domain(15, 25, 63).
        Build()
    
    testStr(t, grad.At(15).Hex(), "#000000")
    testStr(t, grad.At(25).Hex(), "#ff0000")
    testStr(t, grad.At(63).Hex(), "#ffffff")
}

func testStr(t *testing.T, result, expected string) {
    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}
