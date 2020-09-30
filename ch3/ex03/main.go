package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	blueBase      = 255
	redBase       = 0
)

var (
	colorFrom = [3]float64{0, 0, 255}
	colorTo   = [3]float64{255, 0, 0}
)

var max = float64(cells)
var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var d = blueBase/cells + 1

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// strokeColor := rgbHex(linearGradient(float64(i + j)))
			strokeColor := getColor(i, j)
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s;'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, strokeColor)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func getColor(i, j int) string {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := (f(x, y) + 1) / 2 * 100
	return rgbHex(linearGradient(z))
}

func linearGradient(z float64) (uint8, uint8, uint8) {
	d := z / max
	r := math.Abs((1-d)*colorFrom[0] + d*(colorTo[0]))
	g := math.Abs((1-d)*colorFrom[1] + d*(colorTo[1]))
	b := math.Abs((1-d)*colorFrom[2] + d*(colorTo[2]))
	return uint8(r), uint8(g), uint8(b)
}

func rgbHex(r, g, b uint8) string {
	red := fmt.Sprintf("%x", r)
	green := fmt.Sprintf("%x", g)
	blue := fmt.Sprintf("%x", b)
	if len(red) == 1 {
		red = "0" + red
	}
	if len(green) == 1 {
		green = "0" + green
	}
	if len(blue) == 1 {
		blue = "0" + blue
	}
	return fmt.Sprintf("#%s%s%s", red, green, blue)
}
