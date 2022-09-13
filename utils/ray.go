package util

type Ray struct {
	Orig, Dir Vector
}

func (a Ray) At(t float64) Vector {
	return a.Orig.add(a.Dir.multiply(t))
}
