package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

// randFloat returns a random number between [0,1]
func randFloat() float64 {
	return rand.Float64()
}

func clamp(x float64, min float64, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

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

func rayColor(r ray, world hitableList) vec3 {
	rec := &hitRecord{}

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
	samples := 100

	// World
	world := hitableList{}
	world.Add(sphere{vec3{0, 0, -1}, 0.5})
	world.Add(sphere{vec3{0, -100.5, -1}, 100})

	// Camera
	cam := InitCamera(16.0/9.0, 2.0, 1.0, point{0, 0, 0})

	image := fmt.Sprintf("P3\n%v %v\n255\n", imageWidth, imageHeight)

	for j := (imageHeight - 1); j >= 0; j-- {
		fmt.Println("\rScanlines remaining: ", j)
		for i := 0; i < imageWidth; i++ {
			pixelColor := color{0, 0, 0}
			for s := 0; s < samples; s++ {
				// fmt.Println(randFloat())
				u := (float64(i) + randFloat()) / (float64(imageWidth) - 1)
				v := (float64(j) + randFloat()) / (float64(imageHeight) - 1)
				r := cam.getRay(u, v)
				pixelColor = pixelColor.add(rayColor(r, world))

			}

			// c := rayColor(r, world)

			image = fmt.Sprintf("%v%v", image, writeColor(pixelColor, samples))
		}
	}

	out := []byte(image)
	err := os.WriteFile("/Users/scoutdarling-blair/lab/path-tracer/test.ppm", out, 0644)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print("\nDone.\n")

}
