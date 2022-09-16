package main

import "fmt"

func writeColor(pixelColor vec3, samples int) string {

	r := pixelColor.x
	g := pixelColor.y
	b := pixelColor.z

	scale := 1.0 / float64(samples)
	r *= scale
	g *= scale
	b *= scale

	// fmt.Println(r, g, b)

	rgb := fmt.Sprintf("%v %v %v\n",
		uint8(255.999*clamp(r, 0.0, 0.999)),
		uint8(255.999*clamp(g, 0.0, 0.999)),
		uint8(255.999*clamp(b, 0.0, 0.999)))
	return rgb
}
