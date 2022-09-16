package main

import "fmt"

func writeColor(pixelColor vec3) string {
	rgb := fmt.Sprintf("%v %v %v\n",
		uint8(255.999*pixelColor.x),
		uint8(255.999*pixelColor.y),
		uint8(255.999*pixelColor.z))
	return rgb
}
