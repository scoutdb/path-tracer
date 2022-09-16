package main

type ray struct {
	origin, direction vec3
}

func (a ray) at(t float64) vec3 {
	// origin + direction * t
	return a.origin.add(a.direction.scalarMult(t))
}
