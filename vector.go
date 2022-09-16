package main

import "math"

type vec3 struct {
	x, y, z float64
}

// Aliases for vec3
type color = vec3

type point = vec3

// methods
func (a vec3) add(b vec3) vec3 {
	a.x += b.x
	a.y += b.y
	a.z += b.z

	return a
}

func (a vec3) sub(b vec3) vec3 {
	a.x -= b.x
	a.y -= b.y
	a.z -= b.z

	return a
}
func (a vec3) scalarMult(b float64) vec3 {
	a.x *= b
	a.y *= b
	a.z *= b

	return a
}

func (a vec3) length() float64 {
	return math.Sqrt(a.lengthSquared())
}
func (a vec3) lengthSquared() float64 {
	return a.x*a.x + a.y*a.y + a.z*a.z
}

// util functions
func dot(a vec3, b vec3) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func divide(a vec3, t float64) vec3 {
	return a.scalarMult((1 / t))
}

func unitVector(v vec3) vec3 {
	return divide(v, v.length())
}
