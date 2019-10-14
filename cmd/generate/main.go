package main

// This example is based on code from @peterhellberg:
// https://github.com/peterhellberg/gfx/pull/3#issuecomment-491432866

import (
	"github.com/peterhellberg/gfx"
	"github.com/xyproto/burnpal"
)

func createPaletteImage(fn string, palette gfx.Palette) error {
	dst := gfx.NewImage(len(palette), 1)

	for x, c := range palette {
		dst.Set(x, 0, c)
	}

	return gfx.SavePNG(fn, gfx.NewResizedImage(dst, 1120, 96))
}

func createSimplexImage(fn string, palette gfx.Palette) error {
	sn := gfx.NewSimplexNoise(123)
	dst := gfx.NewImage(1120, 96)

	gfx.EachImageVec(dst, gfx.ZV, func(u gfx.Vec) {
		n := sn.Noise2D(u.X/100, u.Y/100)
		c := palette.At((1 + n) / 2)

		gfx.SetVec(dst, u, c)
	})

	return gfx.SavePNG(fn, dst)
}

func main() {
	createPaletteImage("gfx-burn-palette.png", burnpal.GfxPalette())
	createSimplexImage("gfx-burn-simplex.png", burnpal.GfxPalette())
}
