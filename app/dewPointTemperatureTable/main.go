package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/theovassiliou/humidity-golang/humidity"
)

func main() {
	// Define ranges & steps for table
	tempStart := 2.0
	tempEnd := 40.0
	tempStep := 2.0
	rHStart := 20.0
	rHEnd := 95.0
	rHstep := 5.0

	// Initialize a new tabwriter
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.AlignRight|tabwriter.Debug)

	// Print table header
	fmt.Fprintln(w, "Air Temperatur (°C)\tDew Point Temp for a given rel humidity %")
	fmt.Fprintf(w, "----- ")
	for rh := rHStart; rh < rHEnd; rh += rHstep {
		fmt.Fprintf(w, "\t%.0f%%", rh)
	}
	fmt.Fprintln(w, "")

	// Print rows
	for temp := tempStart; temp < tempEnd; temp += tempStep {
		fmt.Fprintf(w, "%.2f", temp)
		for rh := rHStart; rh < rHEnd; rh += rHstep {
			ah := humidity.DewPointTemperature(rh, temp)
			fmt.Fprintf(w, "\t%.2f", ah)
		}
		fmt.Fprintf(w, "\n")
	}

	// Flush to apply tab alignments and write to the standard output
	w.Flush()
}
