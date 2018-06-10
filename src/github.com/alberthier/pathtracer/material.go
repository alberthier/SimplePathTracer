package pathtracer

import "math/rand"

type Material interface {
	Scatter(ray *Ray, record *HitRecord) (attenuation *Vector3, scattered *Ray)
}

func NewMaterial(tp string) Material {
	switch tp {
	case "lambert":
		return &LambertMaterial{Vector3{0.6, 0.1, 0.1}}
		/*case "metal":
			return &Material{}
		case "dielectric":
			return &Material{}*/
	}
	return nil
}

type LambertMaterial struct {
	albedo Vector3
}

func (self *LambertMaterial) Scatter(ray *Ray, record *HitRecord) (attenuation *Vector3, scattered *Ray) {
	var rnd *Vector3
	ok := false
	for !ok {
		rnd = NewVector(rand.Float32(), rand.Float32(), rand.Float32()).Scale(2.0).Substract(vector111)
		ok = rnd.SquaredLength() >= 1.0
	}
	target := record.point.Add(record.normal).Add(rnd)
	scattered = NewRay(record.point, target.Substract(record.point))
	return &self.albedo, scattered
}
