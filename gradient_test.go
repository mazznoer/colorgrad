package colorgrad

import (
	"fmt"
	"testing"
)

func Test_Basic(t *testing.T) {
	test(t, Rgb(1, 0.8431, 0, 1).HexString(), "#ffd700")
	test(t, Rgb8(46, 139, 87, 255).HexString(), "#2e8b57")
	test(t, Hwb(330, 0.4118, 0, 1).HexString(), "#ff69b4")

	test(t, BlendRgb.String(), "BlendRgb")
	test(t, fmt.Sprintf("%s", BlendLinearRgb), "BlendLinearRgb")
	test(t, fmt.Sprintf("%v", BlendOklab), "BlendOklab")

	test(t, InterpolationLinear.String(), "InterpolationLinear")
	test(t, fmt.Sprintf("%s", InterpolationCatmullRom), "InterpolationCatmullRom")
	test(t, fmt.Sprintf("%v", InterpolationBasis), "InterpolationBasis")
}

func Test_GetColors(t *testing.T) {
	grad, _ := NewGradient().Build()
	test(t, len(grad.Colors(0)), 0)
	test(t, grad.Colors(1)[0].HexString(), "#000000")
	testSlice(t, colors2hex(grad.Colors(2)), []string{
		"#000000",
		"#ffffff",
	})
	testSlice(t, colors2hex(grad.Colors(3)), []string{
		"#000000",
		"#808080",
		"#ffffff",
	})

	grad, _ = NewGradient().
		HtmlColors("#f00", "#0f0", "#00f").
		Domain(-1, 1).
		Build()

	testSlice(t, colors2hex(grad.Colors(5)), []string{
		"#ff0000",
		"#808000",
		"#00ff00",
		"#008080",
		"#0000ff",
	})
}

func Test_SpreadRepeat(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("#000", "#fff").
		Build()

	test(t, grad.RepeatAt(-2.0).HexString(), "#000000")
	test(t, grad.RepeatAt(-1.9).HexString(), "#1a1a1a")
	test(t, grad.RepeatAt(-1.5).HexString(), "#808080")
	test(t, grad.RepeatAt(-1.1).HexString(), "#e5e5e5")

	test(t, grad.RepeatAt(-1.0).HexString(), "#000000")
	test(t, grad.RepeatAt(-0.9).HexString(), "#191919")
	test(t, grad.RepeatAt(-0.5).HexString(), "#808080")
	test(t, grad.RepeatAt(-0.1).HexString(), "#e6e6e6")

	test(t, grad.RepeatAt(0.0).HexString(), "#000000")
	test(t, grad.RepeatAt(0.1).HexString(), "#1a1a1a")
	test(t, grad.RepeatAt(0.5).HexString(), "#808080")
	test(t, grad.RepeatAt(0.9).HexString(), "#e5e5e5")

	test(t, grad.RepeatAt(1.0).HexString(), "#000000")
	test(t, grad.RepeatAt(1.1).HexString(), "#1a1a1a")
	test(t, grad.RepeatAt(1.5).HexString(), "#808080")
	test(t, grad.RepeatAt(1.9).HexString(), "#e5e5e5")

	test(t, grad.RepeatAt(2.0).HexString(), "#000000")
	test(t, grad.RepeatAt(2.1).HexString(), "#1a1a1a")
	test(t, grad.RepeatAt(2.5).HexString(), "#808080")
	test(t, grad.RepeatAt(2.9).HexString(), "#e5e5e5")
}

func Test_SpreadReflect(t *testing.T) {
	grad, _ := NewGradient().
		HtmlColors("#000", "#fff").
		Build()

	test(t, grad.ReflectAt(-2.0).HexString(), "#000000")
	test(t, grad.ReflectAt(-1.9).HexString(), "#1a1a1a")
	test(t, grad.ReflectAt(-1.5).HexString(), "#808080")
	test(t, grad.ReflectAt(-1.1).HexString(), "#e5e5e5")

	test(t, grad.ReflectAt(-1.0).HexString(), "#ffffff")
	test(t, grad.ReflectAt(-0.9).HexString(), "#e5e5e5")
	test(t, grad.ReflectAt(-0.5).HexString(), "#808080")
	test(t, grad.ReflectAt(-0.1).HexString(), "#1a1a1a")

	test(t, grad.ReflectAt(0.0).HexString(), "#000000")
	test(t, grad.ReflectAt(0.1).HexString(), "#1a1a1a")
	test(t, grad.ReflectAt(0.5).HexString(), "#808080")
	test(t, grad.ReflectAt(0.9).HexString(), "#e5e5e5")

	test(t, grad.ReflectAt(1.0).HexString(), "#ffffff")
	test(t, grad.ReflectAt(1.1).HexString(), "#e5e5e5")
	test(t, grad.ReflectAt(1.5).HexString(), "#808080")
	test(t, grad.ReflectAt(1.9).HexString(), "#1a1a1a")

	test(t, grad.ReflectAt(2.0).HexString(), "#000000")
	test(t, grad.ReflectAt(2.1).HexString(), "#1a1a1a")
	test(t, grad.ReflectAt(2.5).HexString(), "#808080")
	test(t, grad.ReflectAt(2.9).HexString(), "#e5e5e5")
}
