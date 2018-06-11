package pathtracer

import (
	"math"
	"math/rand"
)

type Material interface {
	Scatter(ray *Ray, record *HitRecord) (attenuation *Vector3, scattered *Ray)
}

func NewMaterial(tp string, diffuse [3]float32, param float32) Material {
	switch tp {
	case "lambert":
		return &LambertMaterial{Vector3{diffuse[0], diffuse[1], diffuse[2]}}
	case "metal":
		return &MetalMaterial{Vector3{diffuse[0], diffuse[1], diffuse[2]}, float32(math.Min(float64(param), 1.0))}
		/*
			case "dielectric":
				return &Material{}*/
	}
	return nil
}

// Lambert =====================================================================

func randomVectorInUnitSphere() *Vector3 {
	var rnd *Vector3
	ok := false
	for !ok {
		rnd = NewVector(rand.Float32(), rand.Float32(), rand.Float32()).Scale(2.0).Substract(vector111)
		ok = rnd.SquaredLength() >= 1.0
	}
	return rnd
}

type LambertMaterial struct {
	albedo Vector3
}

func (self *LambertMaterial) Scatter(ray *Ray, record *HitRecord) (attenuation *Vector3, scattered *Ray) {
	rnd := randomVectorInUnitSphere()
	target := record.point.Add(record.normal).Add(rnd)
	scattered = NewRay(record.point, target.Substract(record.point))
	return &self.albedo, scattered
}

// Metal =====================================================================

type MetalMaterial struct {
	albedo    Vector3
	fuzziness float32
}

func (self *MetalMaterial) Scatter(ray *Ray, record *HitRecord) (attenuation *Vector3, scattered *Ray) {
	v := ray.Direction.Unit()
	reflected := v.Substract(record.normal.Scale(2.0 * v.Dot(record.normal)))
	if reflected.Dot(record.normal) <= 0.0 {
		return nil, nil
	}
	scattered = NewRay(record.point, reflected.Add(randomVectorInUnitSphere().Scale(self.fuzziness)))
	return &self.albedo, scattered
}
