package main

type Camera struct {
	vph, vpw                    float64
	aspect, focus               float64
	origin, horiz, vert, corner vec3
}

func InitCamera(aspectRatio float64, viewportHeight float64, focalLength float64, origin vec3) Camera {
	var c Camera

	c.aspect = aspectRatio
	c.vph = viewportHeight
	c.vpw = aspectRatio * viewportHeight
	c.focus = focalLength

	c.origin = origin
	c.horiz = vec3{c.vpw, 0.0, 0.0}
	c.vert = vec3{0.0, viewportHeight, 0.0}
	c.corner = origin.sub(divide(c.horiz, 2)).
		sub(divide(c.vert, 2)).sub(vec3{0, 0, focalLength})

	return c
}

func (c Camera) getRay(u float64, v float64) ray {
	return ray{c.origin, c.corner.add(c.horiz.scalarMult(u)).add(c.vert.scalarMult(v)).sub(c.origin)}
}
