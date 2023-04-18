package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	points := make(plotter.XYs, 5)
	points[0].X = 0
	points[0].Y = 0
	points[1].X = 1
	points[1].Y = 0
	points[2].X = 1
	points[2].Y = 1
	points[3].X = 0
	points[3].Y = 1
	points[4].X = 0
	points[4].Y = 0

	rect, err := plotter.NewPolygon(points)
	if err != nil {
		panic(err)
	}

	p := plot.New()
	p.Add(rect)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "square.png"); err != nil {
		panic(err)
	}
}
