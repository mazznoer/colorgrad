package colorgrad

func Viridis() Gradient {
	colors := []string{"#440154", "#482777", "#3f4a8a", "#31678e", "#26838f", "#1f9d8a", "#6cce5a", "#b6de2b", "#fee825"}
	grad, _ := NewGradient().
		HexColors(colors...).
		Build()
	return grad
}

func Spectral() Gradient {
	colors := []string{"#9e0142", "#d53e4f", "#f46d43", "#fdae61", "#fee08b", "#ffffbf", "#e6f598", "#abdda4", "#66c2a5", "#3288bd", "#5e4fa2"}
	grad, _ := NewGradient().
		HexColors(colors...).
		Build()
	return grad
}

func Plasma() Gradient {
	colors := []string{"#0d0887", "#42039d", "#6a00a8", "#900da3", "#b12a90", "#cb4678", "#e16462", "#f1834b", "#fca636", "#fccd25", "#f0f921"}
	grad, _ := NewGradient().
		HexColors(colors...).
		Build()
	return grad
}

func Magma() Gradient {
	colors := []string{"#000004", "#140e37", "#3b0f70", "#641a80", "#8c2981", "#b63679", "#de4968", "#f66f5c", "#fe9f6d", "#fece91", "#fcfdbf"}
	grad, _ := NewGradient().
		HexColors(colors...).
		Build()
	return grad
}

func Inferno() Gradient {
	colors := []string{"#000004", "#170b3a", "#420a68", "#6b176e", "#932667", "#bb3654", "#dd513a", "#f3771a", "#fca50a", "#f6d644", "#fcffa4"}
	grad, _ := NewGradient().
		HexColors(colors...).
		Build()
	return grad
}

// Sequential (Single Hue)

func Blues() Gradient {
	colors := []string{"#f7fbff", "#deebf7", "#c6dbef", "#9ecae1", "#6baed6", "#4292c6", "#2171b5", "#08519c", "#08306b"}
	grad, _ := NewGradient().
		HexColors(colors...).
		Build()
	return grad
}

func Greens() Gradient {
	colors := []string{"#f7fcf5", "#e5f5e0", "#c7e9c0", "#a1d99b", "#74c476", "#41ab5d", "#238b45", "#006d2c", "#00441b"}
	grad, _ := NewGradient().
		HexColors(colors...).
		Build()
	return grad
}

func Greys() Gradient {
	grad, _ := NewGradient().
		HexColors("#ffffff", "#000000").
		Build()
	return grad
}

func Oranges() Gradient {
	colors := []string{"#fff5eb", "#fee6ce", "#fdd0a2", "#fdae6b", "#fd8d3c", "#f16913", "#d94801", "#a63603", "#7f2704"}
	grad, _ := NewGradient().
		HexColors(colors...).
		Build()
	return grad
}

func Purples() Gradient {
	colors := []string{"#fcfbfd", "#efedf5", "#dadaeb", "#bcbddc", "#9e9ac8", "#807dba", "#6a51a3", "#54278f", "#3f007d"}
	grad, _ := NewGradient().
		HexColors(colors...).
		Build()
	return grad
}

func Reds() Gradient {
	colors := []string{"#fff5f0", "#fee0d2", "#fcbba1", "#fc9272", "#fb6a4a", "#ef3b2c", "#cb181d", "#a50f15", "#67000d"}
	grad, _ := NewGradient().
		HexColors(colors...).
		Build()
	return grad
}
