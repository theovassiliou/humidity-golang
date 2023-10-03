// ... other imports ...
package main

import (
	"fmt"
	"math"

	"github.com/theovassiliou/humidity-golang/humidity"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// ... rest of the code ...

func main() {
	// Create a new plot
	p := plot.New()

	p.Title.Text = "Absolute Humidity vs Temperature"
	p.X.Label.Text = "Temperature (°C)"
	p.Y.Label.Text = "Absolute Humidity (g/m^3)"

	// Define the temperature range for which you want to plot the graph, e.g., -10°C to 40°C
	tempStart := -10.0
	tempEnd := 40.0

	// Create a points slice with data
	pts := make(plotter.XYs, int(tempEnd-tempStart)+1)
	pts80 := make(plotter.XYs, int(tempEnd-tempStart)+1)
	pts60 := make(plotter.XYs, int(tempEnd-tempStart)+1)
	pts40 := make(plotter.XYs, int(tempEnd-tempStart)+1)
	pts20 := make(plotter.XYs, int(tempEnd-tempStart)+1)

	for i := tempStart; i <= tempEnd; i++ {
		pts[int(i-tempStart)].X = i
		pts80[int(i-tempStart)].X = i
		pts60[int(i-tempStart)].X = i
		pts40[int(i-tempStart)].X = i
		pts20[int(i-tempStart)].X = i
		pts[int(i-tempStart)].Y = humidity.RelativeToAbsolute(100, i)  // Assuming 50% relative humidity for demonstration
		pts80[int(i-tempStart)].Y = humidity.RelativeToAbsolute(80, i) // Assuming 100% relative humidity for demonstration
		pts60[int(i-tempStart)].Y = humidity.RelativeToAbsolute(60, i) // Assuming 100% relative humidity for demonstration
		pts40[int(i-tempStart)].Y = humidity.RelativeToAbsolute(40, i) // Assuming 100% relative humidity for demonstration
		pts20[int(i-tempStart)].Y = humidity.RelativeToAbsolute(20, i) // Assuming 100% relative humidity for demonstration

	}

	// Create a line plotter and set its style
	line, _ := plotter.NewLine(pts)
	line80, _ := plotter.NewLine(pts80)
	line80.LineStyle = plotter.DefaultQuartWhiskerStyle
	line60, _ := plotter.NewLine(pts60)
	line60.LineStyle = plotter.DefaultQuartWhiskerStyle
	line40, _ := plotter.NewLine(pts40)
	line40.LineStyle = plotter.DefaultQuartWhiskerStyle
	line20, _ := plotter.NewLine(pts20)
	line20.LineStyle = plotter.DefaultQuartWhiskerStyle

	p.Add(line)
	p.Add(line80)
	p.Add(line60)
	p.Add(line40)
	p.Add(line20)

	// Adding vertical and horizontal dotted lines
	grid := plotter.NewGrid()
	grid.Horizontal.Color = plotutil.Color(1)                        // Choosing a color, you can adjust this
	grid.Horizontal.Dashes = []vg.Length{vg.Points(2), vg.Points(2)} // Dotted pattern
	grid.Horizontal.Width = vg.Points(0.5)
	grid.Vertical.Color = plotutil.Color(1)                        // Choosing a color, you can adjust this
	grid.Vertical.Dashes = []vg.Length{vg.Points(2), vg.Points(2)} // Dotted pattern
	grid.Vertical.Width = vg.Points(0.5)
	p.Add(grid)

	// Set the X and Y tickers to generate ticks every 10 units away.
	p.X.Tick.Marker = plot.TickerFunc(func(min, max float64) []plot.Tick {
		var ticks []plot.Tick
		for i := math.Floor(min/10) * 10; i <= max; i += 10 {
			if i >= min {
				ticks = append(ticks, plot.Tick{Value: i, Label: fmt.Sprintf("%v", i)})
			}
		}
		return ticks
	})

	p.Y.Tick.Marker = plot.TickerFunc(func(min, max float64) []plot.Tick {
		var ticks []plot.Tick
		for i := math.Floor(min/10) * 10; i <= max; i += 10 {
			if i >= min {
				ticks = append(ticks, plot.Tick{Value: i, Label: fmt.Sprintf("%v", i)})
			}
		}
		return ticks
	})

	labels, _ := plotter.NewLabels(plotter.XYLabels{
		XYs: []plotter.XY{
			{X: tempEnd, Y: humidity.RelativeToAbsolute(100, tempEnd)},
			{X: tempEnd, Y: humidity.RelativeToAbsolute(80, tempEnd)},
			{X: tempEnd, Y: humidity.RelativeToAbsolute(60, tempEnd)},
			{X: tempEnd, Y: humidity.RelativeToAbsolute(40, tempEnd)},
			{X: tempEnd, Y: humidity.RelativeToAbsolute(20, tempEnd)},
		},
		Labels: []string{"100%", "80%", "60%", "40%", "20%"},
	})

	p.Add(labels)
	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "relative_to_absolute.png"); err != nil {
		panic(err)
	}
}
