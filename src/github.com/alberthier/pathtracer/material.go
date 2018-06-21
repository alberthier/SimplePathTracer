package pathtracer

import (
	"math"
	"math/rand"
)

type Material interface {
	Scatter(rng *rand.Rand, ray *Ray, record *HitRecord) (attenuation *Vector3, scattered *Ray)
}

func NewMaterial(tp string, diffuse [3]float64, param float64) Material {
	switch tp {
	case "lambert":
		return &LambertMaterial{Vector3{diffuse[0], diffuse[1], diffuse[2]}}
	case "metal":
		return &MetalMaterial{Vector3{diffuse[0], diffuse[1], diffuse[2]}, math.Min(param, 1.0)}
	case "dielectric":
		return &DielectricMaterial{param}
	}
	return nil
}

func randomVectorInUnitSphere(rng *rand.Rand) *Vector3 {
	for {
		r := NewVector(rng.Float64(), rng.Float64(), rng.Float64())
		p := r.Scale(2.0).Subtract(UnitVector)
		if p.SquaredLength() >= 1.0 {
			return p
		}
	}
}

func reflect(v *Vector3, n *Vector3) *Vector3 {
	return v.Subtract(n.Scale(2.0 * v.Dot(n)))
}

func refract(v *Vector3, n *Vector3, niOverNt float64) *Vector3 {
	uv := v.Unit()
	dot := uv.Dot(n)
	disc := 1.0 - niOverNt*niOverNt*(1.0-dot*dot)
	if disc > 0 {
		sqrtDisc := math.Sqrt(disc)
		return uv.Subtract(n.Scale(dot)).Scale(niOverNt).Subtract(n.Scale(sqrtDisc))
	}
	return nil
}

// Lambert =====================================================================

type LambertMaterial struct {
	albedo Vector3
}

func (self *LambertMaterial) Scatter(rng *rand.Rand, ray *Ray, record *HitRecord) (attenuation *Vector3, scattered *Ray) {
	rnd := randomVectorInUnitSphere(rng)
	target := record.point.Add(record.normal).Add(rnd)
	scattered = NewRay(record.point, target.Subtract(record.point))
	return &self.albedo, scattered
}

// Metal =====================================================================

type MetalMaterial struct {
	albedo    Vector3
	fuzziness float64
}

func (self *MetalMaterial) Scatter(rng *rand.Rand, ray *Ray, record *HitRecord) (attenuation *Vector3, scattered *Ray) {
	reflected := reflect(ray.Direction, record.normal)
	if reflected.Dot(record.normal) <= 0.0 {
		return nil, nil
	}
	scattered = NewRay(record.point, reflected.Add(randomVectorInUnitSphere(rng).Scale(self.fuzziness)))
	return &self.albedo, scattered
}

// Dielectric =====================================================================

type DielectricMaterial struct {
	refractiveIndex float64
}

func schlick(cosine float64, refractiveIndex float64) float64 {
	r0 := (1.0 - refractiveIndex) / (1.0 + refractiveIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1.0-cosine), 5.0)
}

func (self *DielectricMaterial) Scatter(rng *rand.Rand, ray *Ray, record *HitRecord) (attenuation *Vector3, scattered *Ray) {
	var outNormal *Vector3
	var niOverNt float64
	var cosine float64

	dot := ray.Direction.Dot(record.normal)
	if dot > 0 {
		outNormal = record.normal.Scale(-1.0)
		niOverNt = self.refractiveIndex
		cosine = dot / ray.Direction.Length()
	} else {
		outNormal = record.normal
		niOverNt = 1.0 / self.refractiveIndex
		cosine = -dot / ray.Direction.Length()
	}

	refracted := refract(ray.Direction, outNormal, niOverNt)
	if refracted != nil {
		if schlick(cosine, self.refractiveIndex) > rng.Float64() {
			refracted = nil
		}
	}
	if refracted != nil {
		scattered = NewRay(record.point, refracted)
	} else {
		reflected := reflect(ray.Direction, outNormal)
		scattered = NewRay(record.point, reflected)
	}

	attenuation = UnitVector

	return attenuation, scattered
}
