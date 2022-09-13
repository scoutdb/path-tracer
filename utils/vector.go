package util

import "math"

type Vector struct {
	X, Y, Z float64
}

func (a Vector) Sub(b Vector) Vector {
	a.X -= b.X
	a.Y -= b.Y
	a.Z -= b.Z

	return a
}

func (a Vector) Add(b Vector) Vector {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z

	return a
}

func (a Vector) Multiply(b float64) Vector {
	a.X *= b
	a.Y *= b
	a.Z *= b

	return a
}

func (a Vector) Devide(t float64) Vector {
	return a.Multiply(1 / t)
}

func Dot(a Vector, b Vector) float64 {
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
	return math.Sqrt(Dot(a, a))
}
func (a Vector) lengthSquared() float64 {
	return Dot(a, a)
}

// Helpers

// NewVector creates a new Vector instance
func NewVector(x float64, y float64, z float64) Vector {
	return Vector{X: x, Y: y, Z: z}
}
