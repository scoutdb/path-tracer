package util

import "math"

type Vector struct {
	X, Y, Z float64
}

func (a Vector) sub(b Vector) Vector {
	a.X -= b.X
	a.Y -= b.Y
	a.Z -= b.Z

	return a
}

func (a Vector) add(b Vector) Vector {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z

	return a
}

func (a Vector) multiply(b Vector) Vector {
	a.X *= b.X
	a.Y *= b.Y
	a.Z *= b.Z

	return a
}

func (a Vector) devide(b Vector) Vector {
	a.X *= b.X
	a.Y *= b.Y
	a.Z *= b.Z

	return a
}

func (a Vector) dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vector) cross(b Vector) Vector {
	return Vector{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

func (a Vector) length() float64 {
	return math.Sqrt(a.dot(a))
}
func (a Vector) lengthSquared() float64 {
	return a.dot(a)
}
