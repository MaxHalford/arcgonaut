package main

import (
	"bufio"
	"flag"
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/lucasb-eyer/go-colorful"
)

func main() {

	// Command-line input
	var in string
	var height int
	var col1 string
	var col2 string

	flag.StringVar(&in, "f", "", `File to be parsed. The file has to
		have a certain format. Every line should be "a>b>amount".`)
	flag.IntVar(&height, "h", 1980, `Height of the output image. The width
		is deduced from the height.`)
	flag.StringVar(&col1, "c1", "#eeef61", "Start color of the gradient.")
	flag.StringVar(&col2, "c2", "#1e3140", "End color of the gradient.")

	flag.Parse()

	// Parse the arcgo file
	names, counter := OpenArcgoFile(in)

	// Comoute width
	width := float64(height) * math.Phi * 1.2
	// Where the first name is located
	top := float64(height) / 10
	// Where the last name is located
	bottom := float64(height) - top

	// Rescale the counter based on the maximal value
	counter = AssignWidths(counter, 1.0, 20.0)
	// Assign linearly separated coordinates to each name
	coordinates := AssignCoordinates(names, top, bottom)
	// Assign "linearly separated" colors to each name
	colors := AssignColors(names, col1, col2)

	// Initialize the image
	img := image.NewRGBA(image.Rect(0, 0, int(width), height))
	gc := draw2dimg.NewGraphicContext(img)

	// Define the x coordinates for the arcs
	xRight := float64(width) * 1.7 / 3.0
	xLeft := float64(width) * 1.3 / 3.0
	// Graph the arcs
	for a := range counter {
		// Get the color of the name
		gc.SetStrokeColor(colors[a])
		for b := range counter[a] {
			// Set where the arc begins
			y1 := coordinates[a]
			// Set where the arc ends
			y2 := coordinates[b]
			// Set the width of the arc
			z := counter[a][b]
			gc.SetLineWidth(2 * z)
			// Define on what side the arc is
			if y1 < y2 {
				// Right side arc
				Arc(gc, xRight, y1, y2)
			} else {
				// Left side arc
				Arc(gc, xLeft, y1, y2)
			}
		}
	}

	// Set the font for writing the names
	draw2d.SetFontFolder("")
	gc.SetFontData(draw2d.FontData{Name: "coolvetica"})
	gc.SetFillColor(image.Black)
	fontSize := float64(height / (3 * len(names)))
	gc.SetFontSize(fontSize)

	// Write the names
	for _, name := range names {
		x := width/2.0 - (float64(len(name))*fontSize)/3.8
		y := coordinates[name] + fontSize/2 - 2
		gc.FillStringAt(name, x, y)
	}
	// Save to file
	path := strings.Split(in, "/")
	out := strings.Join([]string{strings.Split(path[len(path)-1], ".")[0], "png"}, ".")
	draw2dimg.SaveToPngFile(out, img)

}

// Arc draws a link between two points
func Arc(gc draw2d.GraphicContext, x, y1, y2 float64) {
	// Center
	yc := (y1 + y2) / 2
	// x radius
	xRadius := y2 - y1
	// y radius
	yRadius := y2 - yc
	// Start at -45 degrees
	startAngle := -math.Pi / 2
	// 180 degrees = half a circle
	angle := math.Pi
	gc.ArcTo(x, yc, xRadius, yRadius, startAngle, angle)
	gc.Stroke()
	gc.Fill()
}

// OpenArcgoFile opens a .arcgo file and parses it
func OpenArcgoFile(filename string) ([]string, map[string]map[string]float64) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Create a counter
	counter := make(map[string]map[string]float64)
	// Go through each line of the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Parse the line
		line := strings.Split(scanner.Text(), ">")
		a, b := line[0], line[1]
		amount, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		// Add the elements to the counter
		if _, ok := counter[a]; ok {
			if _, ok := counter[a][b]; ok {
				counter[a][b] += amount
			} else {
				counter[a][b] = amount
			}
		} else {
			counter[a] = make(map[string]float64)
			counter[a][b] = amount
		}
	}
	// Extract the names, they have to be unique
	encountered := make(map[string]bool)
	names := []string{}
	for a := range counter {
		if encountered[a] == true {
			// pass
		} else {
			encountered[a] = true
			names = append(names, a)
		}
		for b := range counter[a] {
			if encountered[b] == true {
				// pass
			} else {
				encountered[b] = true
				names = append(names, b)
			}
		}
	}
	return names, counter
}

// AssignCoordinates assigns a y value to each name.
// The y values are linearly separated between the top
// and the bottom.
func AssignCoordinates(names []string, min, max float64) map[string]float64 {
	sort.Strings(names)
	length := len(names)
	step := (max - min) / float64(length-1)
	coordinates := make(map[string]float64)
	for i, name := range names {
		coordinates[name] = min + float64(i)*step
	}
	return coordinates
}

// AssignColors assigns an RGB color to each name.
// The colors are generated with the colorful library.
func AssignColors(names []string, minHex, maxHex string) map[string]color.Color {
	c1, _ := colorful.Hex(minHex)
	c2, _ := colorful.Hex(maxHex)
	length := len(names)
	colors := make(map[string]color.Color)
	for i, name := range names {
		c := c1.BlendHsv(c2, float64(i)/float64(length-1))
		m := float64(150)
		rgb := color.RGBA{uint8(m * c.R), uint8(m * c.G), uint8(m * c.B), uint8(m)}
		colors[name] = rgb
	}
	return colors
}

// AssignWidths rescales the counts. This allows
// user values to be negative or/and huge. All the
// values are rescaled in a defined interval so
// that the brush strokes are not too big/invalid.
func AssignWidths(counter map[string]map[string]float64, newMin, newMax float64) map[string]map[string]float64 {
	// Search for the minimal and the maximal count values
	min := math.Inf(1)
	max := math.Inf(-1)
	for a := range counter {
		for b := range counter[a] {
			// Check for new min
			if counter[a][b] < min {
				min = counter[a][b]
			}
			// Check for new max
			if counter[a][b] > max {
				max = counter[a][b]
			}
		}
	}
	// Rescale every value
	for a := range counter {
		for b := range counter[a] {
			// Check for new min
			normalized := (counter[a][b] - min) / (max - min)
			rescaled := normalized*(newMax-newMin) + newMin
			counter[a][b] = rescaled
		}
	}
	return counter
}
