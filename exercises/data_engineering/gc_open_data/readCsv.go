package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const timeLayout = "2006-01-02 15:04:05"

func readGCCsv(fullfilepath string) {
	fileBase := strings.Split(filepath.Base(fullfilepath), ".")[0]

	file, err := os.Open(fullfilepath)

	if err != nil {
		fmt.Printf("Error occured: %s", err)
	}
	defer file.Close()

	if err != nil {
		fmt.Println("Error occured during read", err)
	}

	parts := strings.Split(fileBase, "_")
	dateStr := fmt.Sprintf("%s-%s-%sT%s:%s:%sZ", parts[0], parts[1], parts[2], parts[3], parts[4], parts[5])
	startTime, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		fmt.Println("Error parsing date", err)
		return
	}

	df := dataframe.ReadCSV(file).Select([]string{"secs", "km", "alt"})

	// Create time index from the start timestamp and delta seconds
	secs := df.Col("secs").Float()
	timeIndex := make([]time.Time, len(secs))
	for i, sec := range secs {
		timeIndex[i] = startTime.Add(time.Duration(sec) * time.Second)
	}

	timeIndexString := make([]string, len(timeIndex))

	for i, t := range timeIndex {
		timeIndexString[i] = t.Format("2006-01-02 15:04:05")
	}

	// fmt.Println(timeIndexString)

	df = df.Mutate(series.New(timeIndexString, series.String, "time_index"))

	df = df.Drop("secs")

	fmt.Println(df)

	log.Println("Simulated DataFrame:")
	log.Println(df)

	// === Part 2: Create the plot from the dataframe ===

	// 1. Extract the "alt" and "time_index" columns
	altitudes := df.Col("alt").Float()
	timeIndexStrings := df.Col("time_index").Records()

	// 2. Prepare data for plotting. We need plotter.XYs which uses float64 for X and Y.
	pts := make(plotter.XYs, df.Nrow())
	for i := 0; i < df.Nrow(); i++ {
		// Parse the time string back into a time.Time object
		t, err := time.Parse(timeLayout, timeIndexStrings[i])
		if err != nil {
			log.Fatalf("Error parsing time string: %v", err)
		}

		// Use the Unix timestamp (a float64) for the X-value
		pts[i].X = float64(t.Unix())
		// Use the altitude for the Y-value
		pts[i].Y = altitudes[i]
	}

	// 3. Create a new plot
	p := plot.New()
	p.Title.Text = "Altitude over Time"
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Altitude (m)"

	// 4. Use a time-aware marker for the X-axis ticks
	// This tells the plot to interpret the float64 X-values as Unix timestamps
	p.X.Tick.Marker = plot.TimeTicks{Format: "15:04:05"} // Format as HH:MM:SS

	// 5. Add the data as a line and add a grid
	line, err := plotter.NewLine(pts)
	if err != nil {
		log.Fatal(err)
	}
	p.Add(line)
	p.Add(plotter.NewGrid())

	// 6. Save the plot to a file
	if err := p.Save(8*vg.Inch, 4*vg.Inch, fileBase+".png"); err != nil {
		log.Fatal(err)
	}

	log.Printf("Plot successfully saved to %s_plot.png", fileBase)
}
