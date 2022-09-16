package main

import "math"

type sphere struct {
	center vec3
	radius float64
}

func (s sphere) hit(r ray, tMin float64, tMax float64, rec *hitRecord) bool {
	oc := r.origin.sub(s.center)
	a := r.direction.lengthSquared()
	halfB := dot(oc, r.direction)
	c := oc.lengthSquared() - s.radius*s.radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false
	}
	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (-halfB - sqrtd) / a
	if root < tMin || tMax < root {
		root = (-halfB + sqrtd) / a
		if root < tMin || tMax < root {
			return false
		}
	}

	rec.t = root
	rec.p = r.at(rec.t)
	outwardNormal := vec3(divide(rec.p.sub(s.center), s.radius))
	rec.setFaceNormal(r, outwardNormal)
	return true
}
