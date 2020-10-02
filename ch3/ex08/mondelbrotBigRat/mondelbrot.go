package mondelbrotBigRat

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
)

func Mondelbrot() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 100
	const contrast = 15
	zReal := (&big.Rat{}).SetFloat64(real(z))
	zImag := (&big.Rat{}).SetFloat64(imag(z))
	var vReal, vImag = &big.Rat{}, &big.Rat{}

	for n := uint8(0); n < iterations; n++ {
		vReal2, vImag2 := &big.Rat{}, &big.Rat{}
		vReal2.Mul(vReal, vReal).Sub(vReal2, (&big.Rat{}).Mul(vImag, vImag)).Add(vReal2, vReal)
		vImag2.Mul(vReal, vImag).Mul(vImag2, big.NewRat(2, 1)).Add(vImag2, zImag)
		vReal, vImag = vReal2.Add(vReal2, zReal), vImag2.Add(vImag2, zImag)
		squareSum := &big.Rat{}
		squareSum.Mul(vReal, vReal).Add(squareSum, (&big.Rat{}).Mul(vImag, vImag))
		if squareSum.Cmp(big.NewRat(2*2, 1)) == 1 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
