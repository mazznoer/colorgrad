package colorgrad

import (
    "image/color"
    "math"
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

    testStr(t, grad.At(math.NaN()).Hex(), "#000000")

    grad, _ = NewGradient().
        Colors(
            color.RGBA{255,0,0,255},
            color.RGBA{255,255,0,255},
            color.RGBA{0,0,255,255},
        ).
        Build()
    colors = grad.Colors(3)

    if len(colors) != 3 {
        t.Errorf("Expected 3, got %v", len(colors))
    }
    testStr(t, colors[0].Hex(), "#ff0000")
    testStr(t, colors[1].Hex(), "#ffff00")
    testStr(t, colors[2].Hex(), "#0000ff")

    testStr(t, grad.At(math.NaN()).Hex(), "#ff0000")
}

func TestDomain(t *testing.T) {
    grad, _ := NewGradient().
        HexColors("#00ff00", "#ff0000", "#ffff00").
        Domain(0, 0.75, 1).
        Build()
    
    testStr(t, grad.At(0).Hex(), "#00ff00")
    testStr(t, grad.At(0.75).Hex(), "#ff0000")
    testStr(t, grad.At(1).Hex(), "#ffff00")

    // outside domain
    testStr(t, grad.At(-1).Hex(), "#00ff00")
    testStr(t, grad.At(1.5).Hex(), "#ffff00")

    grad, _ = NewGradient().
        HexColors("#00ff00", "#ff0000", "#0000ff", "#ffff00").
        Domain(15, 25, 29, 63).
        Build()
    
    testStr(t, grad.At(15).Hex(), "#00ff00")
    testStr(t, grad.At(25).Hex(), "#ff0000")
    testStr(t, grad.At(29).Hex(), "#0000ff")
    testStr(t, grad.At(63).Hex(), "#ffff00")

    // outside domain
    testStr(t, grad.At(10).Hex(), "#00ff00")
    testStr(t, grad.At(67).Hex(), "#ffff00")
}

func testStr(t *testing.T, result, expected string) {
    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}
