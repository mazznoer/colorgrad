package colorgrad

import (
	"testing"
)

func TestPresetGradients(t *testing.T) {
	var grad Gradient

	grad = CubehelixDefault()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#000000", "#19122c", "#1b354c", "#2c5c48", "#3f7533", "#7e7a36", "#bc7967", "#d486b0", "#cba8e6", "#c1d2f3", "#ddf0ef", "#ffffff",
	})

	grad = Warm()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#6e40aa", "#923db3", "#b83cb0", "#da3fa3", "#f6478d", "#ff5572", "#ff6956", "#ff823e", "#f59f30", "#ddbd30", "#c4d93e", "#aff05b",
	})

	grad = Cool()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#6e40aa", "#6252c5", "#5069d9", "#3c84e1", "#42a0dd", "#49bbcd", "#51d3b5", "#5ae597", "#65f17a", "#71f663", "#83f557", "#aff05b",
	})

	grad = Rainbow()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#6e40aa", "#b83cb0", "#f6478d", "#ff6956", "#f59f30", "#c4d93e", "#83f557", "#65f17a", "#51d3b5", "#42a0dd", "#5069d9", "#6e40aa",
	})

	grad = Cividis()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#002051", "#083069", "#24416e", "#44516d", "#5f626e", "#757372", "#898477", "#9d9778", "#b4aa73", "#d0be67", "#ecd354", "#fdea45",
	})

	grad = Sinebow()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#ff4040", "#eb860e", "#b4c901", "#6df61b", "#2cfd56", "#05dc9e", "#059edc", "#2c56fd", "#6d1bf6", "#b401c9", "#eb0e86", "#ff4040",
	})

	grad = Turbo()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#23171b", "#4a51d4", "#3491f8", "#25c9d5", "#3aef9a", "#71fe65", "#b8f140", "#f2cb2c", "#ff9220", "#ed5215", "#b41d07", "#900c00",
	})

	grad = Viridis()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#440154", "#461c6c", "#43377f", "#3c4f89", "#33648d", "#2a798e", "#248d8d", "#31a480", "#5ec263", "#94d641", "#cae02c", "#fee825",
	})

	grad = Plasma()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#0d0887", "#3c049a", "#6302a5", "#850ba3", "#a52097", "#bf3983", "#d5546e", "#e76f5a", "#f48d45", "#fbad33", "#f9d226", "#f0f921",
	})

	grad = Magma()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#000004", "#150b33", "#341062", "#59177b", "#7e2380", "#a3307c", "#c83f71", "#e65864", "#f77d63", "#fda775", "#fed296", "#fcfdbf",
	})

	grad = Inferno()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#000004", "#170834", "#3a0b5c", "#60146b", "#842169", "#a92f5c", "#ca4348", "#e45f2e", "#f58417", "#faae1b", "#f8d951", "#fcffa4",
	})

	grad = BrBG()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#543005", "#86500d", "#b47a2b", "#d6af67", "#edd9a9", "#f4eedc", "#deefec", "#acded7", "#6bbdb2", "#2e8f86", "#07635a", "#003c30",
	})

	grad = PRGn()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#40004b", "#6f2a7c", "#9362a3", "#b796c4", "#d9c2de", "#eee6ef", "#e8f2e5", "#c5e8c0", "#90cd8e", "#50a35b", "#1d7436", "#00441b",
	})

	grad = PiYG()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#8e0152", "#bc217a", "#d964a5", "#eba3cd", "#f8d0e7", "#f9ecf2", "#eff5e3", "#d4edb4", "#a8d674", "#77b43f", "#4b8d23", "#276419",
	})

	grad = PuOr()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#2d004b", "#51287f", "#7963a6", "#a49bc7", "#cac8e1", "#e8e8ef", "#f9ebd7", "#fdd197", "#f3a84e", "#d67b17", "#ad5708", "#7f3b08",
	})

	grad = RdBu()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#67001f", "#a61c2d", "#cf5349", "#ea9175", "#f9c6ad", "#f9e9df", "#e4edf2", "#b9d9e9", "#7cb6d6", "#418bbf", "#1f609f", "#053061",
	})

	grad = RdGy()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#67001f", "#a61c2d", "#cf5349", "#ea9175", "#f9c6ad", "#fdede3", "#f0efee", "#d2d2d2", "#ababab", "#7c7c7c", "#494949", "#1a1a1a",
	})

	grad = RdYlBu()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#a50026", "#d02d2a", "#ed623e", "#fa9b5a", "#fecd7f", "#feefaa", "#f0f8d8", "#cce9ef", "#9ccce2", "#6ca2cb", "#476eb1", "#313695",
	})

	grad = RdYlGn()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#a50026", "#d02d2a", "#ed623e", "#fa9b5a", "#fecd7c", "#fdefa5", "#ecf6a5", "#c6e780", "#94d16a", "#57b55e", "#1e924d", "#006837",
	})

	grad = Spectral()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#9e0142", "#cd374b", "#ec6649", "#fa9b5a", "#fecd7c", "#fef0a5", "#f2f9ac", "#cfec9e", "#98d5a4", "#5eb5ab", "#4283b4", "#5e4fa2",
	})

	grad = Blues()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#f7fbff", "#e5eff9", "#d3e4f3", "#bdd8ec", "#a0cae3", "#7eb8da", "#5da4d0", "#408ec4", "#2877b7", "#1460a7", "#0a488d", "#08306b",
	})

	grad = Greens()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#f7fcf5", "#e9f7e5", "#d7efd1", "#bfe6b9", "#a4da9e", "#84cb84", "#61bb6d", "#41a75b", "#289149", "#117a38", "#026128", "#00441b",
	})

	grad = Greys()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#ffffff", "#f4f4f4", "#e5e5e5", "#d3d3d3", "#bebebe", "#a4a4a4", "#898989", "#707070", "#575757", "#393939", "#1b1b1b", "#000000",
	})

	grad = Oranges()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#fff5eb", "#feead5", "#fedcb9", "#fdc997", "#fdb171", "#fc994d", "#f8802e", "#ed6614", "#db4f06", "#bd3e02", "#9c3203", "#7f2704",
	})

	grad = Purples()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#fcfbfd", "#f2f0f7", "#e5e4f0", "#d4d4e8", "#bfbfdd", "#a9a7cf", "#9390c3", "#7f77b7", "#6e59a7", "#5e3a98", "#4e1d8a", "#3f007d",
	})

	grad = Reds()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#fff5f0", "#fee5d9", "#fdcfbb", "#fcb399", "#fc9677", "#fb7859", "#f65940", "#e9392d", "#d12120", "#b61319", "#930b13", "#67000d",
	})

	grad = BuGn()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#f7fcfd", "#e9f7f9", "#d9f1f0", "#c0e7e0", "#9edacb", "#79cab1", "#59bb93", "#3fa971", "#289250", "#117a38", "#026128", "#00441b",
	})

	grad = BuPu()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#f7fcfd", "#e6f0f6", "#d1e0ee", "#b9cfe4", "#a3bcda", "#93a3cd", "#8d86be", "#8b67af", "#88489f", "#83278a", "#700d6e", "#4d004b",
	})

	grad = GnBu()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#f7fcf0", "#e6f6e1", "#d7efd1", "#c4e8c3", "#aadeba", "#8bd2bf", "#6bc2c9", "#4caecd", "#3193c2", "#1978b4", "#0a5d9f", "#084081",
	})

	grad = OrRd()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#fff7ec", "#feecd1", "#fedfb5", "#fdcf9b", "#fdbb84", "#fc9e6a", "#f77f54", "#eb5f41", "#da3a27", "#c3170f", "#a40302", "#7f0000",
	})

	grad = PuBuGn()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#fff7fb", "#f1e8f3", "#dfdaeb", "#c7cde4", "#a7bfdc", "#7eb0d3", "#56a0c9", "#3190b6", "#108394", "#027570", "#016150", "#014636",
	})

	grad = PuBu()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#fff7fb", "#f1ebf4", "#dfddec", "#c7cee4", "#a9bfdc", "#86b0d3", "#5da0c9", "#338cbe", "#1277b1", "#05649c", "#03507d", "#023858",
	})

	grad = PuRd()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#f7f4f9", "#ebe5f1", "#decee5", "#d3b3d7", "#ce96c8", "#d775b8", "#e14fa1", "#e12c84", "#d01762", "#b0094c", "#8b0138", "#67001f",
	})

	grad = RdPu()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#fff7f3", "#fee6e3", "#fdd3d0", "#fcbdc0", "#faa0b5", "#f77ca9", "#ec559e", "#d62f93", "#b60f84", "#92027a", "#6d0173", "#49006a",
	})

	grad = YlGnBu()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#ffffd9", "#f1f9bf", "#dbf1b4", "#b7e3b6", "#87d0bb", "#59bec0", "#35a8c2", "#238bbb", "#2168ad", "#23489c", "#1b2f81", "#081d58",
	})

	grad = YlGn()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#ffffe5", "#f8fcc6", "#e9f6b0", "#d0ec9f", "#b0de90", "#8bce80", "#64bc6f", "#41a65b", "#288c49", "#11753d", "#025e33", "#004529",
	})

	grad = YlOrBr()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#ffffe5", "#fff8c7", "#ffeda8", "#fedc83", "#fec559", "#fda938", "#f78a22", "#e76d13", "#d05407", "#b03f03", "#8b3005", "#662506",
	})

	grad = YlOrRd()
	testSlice(t, colors2hex(grad.Colors(12)), []string{
		"#ffffcc", "#fff2ac", "#ffe48d", "#fed06e", "#feb653", "#fd9942", "#fc7535", "#f74b29", "#e62621", "#cd0d22", "#ab0225", "#800026",
	})

	// Cyclical gradients

	grad = Rainbow()
	test(t, grad.At(0).HexString(), grad.At(1).HexString())

	grad = Sinebow()
	test(t, grad.At(0).HexString(), grad.At(1).HexString())
}
