package main

import (
	"fmt"
	"os"

	util "github.com/scoutdb/path-tracer/utils"
)

const image_width = 4
const image_height = 4

type Output struct {
	Image string
}

func writeColor(pixelColor util.Vector) string {
	ty := fmt.Sprintf("%v %v %v\n",
		int(255.999*pixelColor.X),
		int(255.999*pixelColor.Y),
		int(255.999*pixelColor.Z))
	return ty
}

func main() {

	// final output string
	ppm := &Output{
		// start the PPM with some formatt info
		// P3 for ASCII colors,
		// Columns (image_width) , Rows (image_height)
		Image: fmt.Sprintf("P3\n%v %v\n255\n", image_width, image_height),
	}

	// Loop through pixel grid
	for j := (image_height - 1); j >= 0; j-- {
		fmt.Println("\rScanlines remaining: ", j)
		for i := 0; i < image_width; i++ {
			r := float64(i) / (image_width - 1)
			g := float64(j) / (image_height - 1)
			b := 0.25

			color := util.Vector{
				X: r,
				Y: g,
				Z: b,
			}

			column := fmt.Sprintf("%v%v", ppm.Image, writeColor(color))

			ppm = &Output{
				Image: column,
			}
		}
	}

	out := []byte(ppm.Image)
	err := os.WriteFile("/Users/scoutdarling-blair/lab/path-tracer/test.ppm", out, 0644)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print("\nDone.\n")
}
