package main

import (
	"fmt"
	"os"

	util "github.com/scoutdb/path-tracer/utils"
)

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

func hitSphere(center util.Vector, radious float64, r util.Ray) bool {
	oc := r.Orig.Sub(center)
	a := util.Dot(r.Dir, r.Dir)
	b := 2.0 * util.Dot(oc, r.Dir)
	c := util.Dot(oc, oc) - radious*radious
	Discriminate := b*b - 4*a*c
	return (Discriminate > 0)
}
func rayColor(r util.Ray) util.Vector {

	hs := hitSphere(util.NewVector(0, 0, -1), 0.5, r)
	// fmt.Println(hs)
	if hs == true {
		return util.Vector{X: 1, Y: 0, Z: 0}
	}

	unitDirection := util.Vector(r.Dir)
	t := 0.5 * (unitDirection.Y + 1)

	C := util.NewVector(1.0, 1.0, 1.0)
	C2 := util.NewVector(0.5, 0.7, 1.0)
	return C.Multiply(1.0 - t).Add(C2.Multiply(t))
}

func main() {

	// image
	const aspectRatio = 16.0 / 9.0
	const imageWidth = 400
	const imageHeight = int(imageWidth / aspectRatio)

	//camera
	const viewportHeight = 2.0
	const viewportWidth = aspectRatio * viewportHeight
	const focalLegnth = 1.0

	origin := util.NewVector(0, 0, 0)
	horizontal := util.NewVector(viewportWidth, 0, 0)
	vertical := util.NewVector(0, viewportHeight, 0)
	lowerLeftCorner := origin.Sub(horizontal.Devide(2)).
		Sub(vertical.Devide(2).Sub(util.NewVector(0, 0, focalLegnth)))

	// final output string
	ppm := &Output{
		// start the PPM with some formatt info
		// P3 for ASCII colors,
		// Columns (image_width) , Rows (image_height)
		Image: fmt.Sprintf("P3\n%v %v\n255\n", imageWidth, imageHeight),
	}

	// Loop through pixel grid
	for j := (imageHeight - 1); j >= 0; j-- {
		fmt.Println("\rScanlines remaining: ", j)
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / (imageWidth - 1)
			v := float64(j) / float64(imageHeight-1)
			// b := 0.25

			r := util.Ray{
				Orig: origin,
				Dir:  lowerLeftCorner.Add(horizontal.Multiply(u).Add(vertical.Multiply(v).Sub(origin))),
			}

			t := rayColor(r)

			column := fmt.Sprintf("%v%v", ppm.Image, writeColor(t))

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
