package surface

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canva size in pixels
	cells         = 100                 // number of prid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30degree)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30d), cos(30d)

func Surface(sgn string) {
	surfaceFunc := chooseSurfaceGraphFunc(sgn)

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, surfaceFunc)
			bx, by := corner(i, j, surfaceFunc)
			cx, cy := corner(i, j+1, surfaceFunc)
			dx, dy := corner(i+1, j+1, surfaceFunc)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Printf("</svg>")
}

func corner(i, j int, sgf surfaceGraphFunc) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := sgf.f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}
