package main

import (
	"fmt"
	"math"
	"os"
)

func hitSphere(center vec3, radius float64, r ray) float64 {
	oc := r.origin.sub(center)
	a := r.direction.lengthSquared()
	halfB := dot(oc, r.direction)
	c := oc.lengthSquared() - radius*radius
	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return -1.0
	} else {
		return ((-halfB - math.Sqrt(discriminant)) / a)
	}
}

// func rayColor(r ray) vec3 {
// 	var hs = hitSphere(point{0, 0, -1}, 0.5, r)
// 	if hs > 0.0 {
// 		N := unitVector(r.at(hs).sub(vec3{0, 0, -1}))
// 		return color{N.x + 1, N.y + 1, N.z + 1}.scalarMult(0.5)
// 	}

// 	unitDirection := unitVector(r.direction)
// 	t := 0.5 * (unitDirection.y + 1.0)

//		white := color{1, 1, 1}.scalarMult(1 - t)
//		blue := color{0.5, 0.7, 1}.scalarMult(t)
//		return white.add(blue)
//	}
func rayColor(r ray, world hitableList) vec3 {
	rec := &hitRecord{}
	// var hs = hitSphere(point{0, 0, -1}, 0.5, r)
	if world.hit(r, 0, 10000, rec) {
		return color{1, 1, 1}.add(rec.normal).scalarMult(0.5)
	}

	unitDirection := unitVector(r.direction)
	t := 0.5 * (unitDirection.y + 1.0)

	white := color{1, 1, 1}.scalarMult(1 - t)
	blue := color{0.5, 0.7, 1}.scalarMult(t)
	return white.add(blue)
}

func main() {

	// Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / (aspectRatio))

	// World
	world := hitableList{}
	world.Add(sphere{vec3{0, 0, -1}, 0.5})
	world.Add(sphere{vec3{0, -100.5, -1}, 100})

	// Camera
	viewportHeight := 2.0
	viewportWidth := float64(aspectRatio * viewportHeight)
	focalLength := 1.0

	origin := point{0, 0, 0}
	horizontal := vec3{viewportWidth, 0, 0}
	vertical := vec3{0, viewportHeight, 0}
	lowerLeftCorner := origin.sub(divide(horizontal, 2)).
		sub(divide(vertical, 2)).sub(vec3{0, 0, focalLength})

	image := fmt.Sprintf("P3\n%v %v\n255\n", imageWidth, imageHeight)

	for j := (imageHeight - 1); j >= 0; j-- {
		fmt.Println("\rScanlines remaining: ", j)
		for i := 0; i < imageWidth; i++ {

			u := float64(i) / (float64(imageWidth) - 1)
			v := float64(j) / (float64(imageHeight) - 1)

			r := ray{origin, lowerLeftCorner.add(horizontal.scalarMult(u)).
				add(vertical.scalarMult(v)).sub(origin)}

			c := rayColor(r, world)

			image = fmt.Sprintf("%v%v", image, writeColor(c))
		}
	}

	out := []byte(image)
	err := os.WriteFile("/Users/scoutdarling-blair/lab/path-tracer/test.ppm", out, 0644)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print("\nDone.\n")

}
