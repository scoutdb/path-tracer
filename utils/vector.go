package main

import "math"

type Vector struct {
	x, y, z float64
}

func (a Vector) sub(b Vector) Vector {
	a.x -= b.x
	a.y -= b.y
	a.z -= b.z

	return a
}

func (a Vector) add(b Vector) Vector {
	a.x += b.x
	a.y += b.y
	a.z += b.z

	return a
}

func (a Vector) multiply(b Vector) Vector {
	a.x *= b.x
	a.y *= b.y
	a.z *= b.z

	return a
}

func (a Vector) devide(b Vector) Vector {
	a.x *= b.x
	a.y *= b.y
	a.z *= b.z

	return a
}

func (a Vector) dot(b Vector) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func (a Vector) cross(b Vector) Vector {
	return Vector{
		x: a.y*b.z - a.z*b.y,
		y: a.z*b.x - a.x*b.z,
		z: a.x*b.y - a.y*b.x,
	}
}

func (a Vector) length() float64 {
	return math.Sqrt(a.dot(a))
}
func (a Vector) lengthSquared() float64 {
	return a.dot(a)
}
