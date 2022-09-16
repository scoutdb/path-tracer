package main

type hitRecord struct {
	p         point
	normal    vec3
	t         float64
	frontface bool
}

type hitableList struct {
	objects []sphere
}

// func hit (r ray, tMin float64, tMax float64, rec hitRecord )bool{
// 	return false
// }

func (l *hitableList) Add(o ...sphere) int {
	l.objects = append(l.objects, o...)
	return len(l.objects)
}

func (l *hitableList) hit(r ray, tMin float64, tMax float64, rec *hitRecord) bool {

	tempRec := rec
	hitAnything := false
	closestSoFar := tMax

	for _, h := range l.objects {
		if h.hit(r, tMin, closestSoFar, tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
		}
	}

	return hitAnything
}

func (h *hitRecord) setFaceNormal(r ray, outwardNormal vec3) {
	frontFace := dot(r.direction, outwardNormal) < 0
	if frontFace {
		h.normal = outwardNormal
	} else {
		h.normal = vec3{-outwardNormal.x, -outwardNormal.y, -outwardNormal.z}
	}
}
