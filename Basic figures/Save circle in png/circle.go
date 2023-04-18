package main

import (
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette/moreland"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	n := 50
	circle := make(plotter.XYs, n)

	for i := range circle {
		theta := float64(i) / float64(n) * 2 * math.Pi
		circle[i].X = math.Cos(theta)
		circle[i].Y = math.Sin(theta)
	}

	p := plot.New()

	scatter, err := plotter.NewScatter(circle)
	if err != nil {
		panic(err)
	}
	color, _ := moreland.ExtendedBlackBody().At(0.5)
	scatter.GlyphStyle.Color = color

	p.Add(scatter)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "circle.png"); err != nil {
		panic(err)
	}
}
