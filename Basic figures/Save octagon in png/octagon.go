package main

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	n := 8
	star := make(plotter.XYs, n*2)

	for i := range star {
		r := 0.5
		if i%2 == 0 {
			r = 0.2
		}
		theta := float64(i) / float64(n) * 2 * math.Pi
		star[i].X = r * math.Cos(theta)
		star[i].Y = r * math.Sin(theta)
	}

	octagon, err := plotter.NewPolygon(star)
	if err != nil {
		panic(err)
	}

	p := plot.New()

	p.Add(octagon)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "octagon.png"); err != nil {
		panic(err)
	}
}
