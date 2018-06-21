package pathtracer

import "math"

type HitRecord struct {
	t      float64
	point  *Vector3
	normal *Vector3
	object SceneObject
}

type SceneObject interface {
	HitBy(ray *Ray, tmin float64, tmax float64, record *HitRecord) bool
	GetMaterial() Material
	Update(t float64)
}

type ObjectBase struct {
	Position AnimatedVector
}

type Sphere struct {
	ObjectBase
	Radius   AnimatedValue
	Material Material
}

func NewSphere(position AnimatedVector, radius AnimatedValue, material Material) *Sphere {
	return &Sphere{
		ObjectBase{position},
		radius, material}
}

func (self *Sphere) HitBy(ray *Ray, tmin float64, tmax float64, record *HitRecord) bool {
	oc := ray.Origin.Subtract(self.Position.Get())
	a := ray.Direction.Dot(ray.Direction)
	b := oc.Dot(ray.Direction)
	radius := self.Radius.Get()
	c := oc.Dot(oc) - radius*radius
	disc := b*b - a*c
	if disc > 0 {
		sd := math.Sqrt(disc)
		t := (-b - sd) / a
		if t <= tmin || tmax <= t {
			t = (-b + sd) / a
			if t <= tmin || tmax <= t {
				return false
			}
		}
		record.t = t
		record.point = ray.PointAt(t)
		record.normal = record.point.Subtract(self.Position.Get()).Scale(1.0 / radius)
		record.object = self
		return true
	}
	return false
}

func (self *Sphere) GetMaterial() Material {
	return self.Material
}

func (self *Sphere) Update(t float64) {
	self.Position.Update(t)
	self.Radius.Update(t)
}
