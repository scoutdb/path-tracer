package main

import (
	"fmt"
	"os"
)

const image_width = 255
const image_height = 255

type Test struct {
	Test string
}

func main() {

	ppm := &Test{}

	tt := fmt.Sprintf("P3\n%v %v\n255\n", image_width, image_height)
	ppm = &Test{
		Test: tt,
	}
	// fmt.Print(tt)

	for j := (image_height - 1); j >= 0; j-- {
		fmt.Println("\rScanlines remaining: ", j)
		for i := 0; i < image_width; i++ {
			r := float64(i) / (image_width - 1)
			g := float64(j) / (image_height - 1)
			b := 0.25

			ir := (255.999 * float64(r))
			ig := (255.999 * float64(g))
			ib := (255.999 * float64(b))

			tt := fmt.Sprintf("%v%v %v %v\n", ppm.Test, int(ir), int(ig), int(ib))

			ppm = &Test{
				Test: tt,
			}
		}
	}

	out := []byte(ppm.Test)
	err := os.WriteFile("/Users/scoutdarling-blair/lab/path-tracer/test.ppm", out, 0644)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print("\nDone.\n")
}
