package pathtracer

import "math"

type HitRecord struct {
	t      float32
	point  *Vector3
	normal *Vector3
	object SceneObject
}

type SceneObject interface {
	HitBy(ray *Ray, tmin float32, tmax float32, record *HitRecord) bool
	GetMaterial() Material
}

type ObjectBase struct {
	Position Vector3
}

type Sphere struct {
	ObjectBase
	Radius   float32
	Material Material
}

func NewSphere(x float32, y float32, z float32, radius float32, material Material) *Sphere {
	return &Sphere{
		ObjectBase{Vector3{x, y, z}},
		radius, material}
}

func (self *Sphere) HitBy(ray *Ray, tmin float32, tmax float32, record *HitRecord) bool {
	oc := ray.Origin.Substract(&self.Position)
	a := ray.Direction.Dot(ray.Direction)
	b := oc.Dot(ray.Direction)
	c := oc.Dot(oc) - self.Radius*self.Radius
	disc := b*b - a*c
	if disc > 0 {
		sd := float32(math.Sqrt(float64(disc)))
		t := (-b - sd) / a
		if t <= tmin || tmax <= t {
			t = (-b + sd) / a
			if t <= tmin || tmax <= t {
				return false
			}
		}
		record.t = t
		record.point = ray.PointAt(t)
		record.normal = record.point.Substract(&self.Position).Scale(1.0 / self.Radius)
		record.object = self
		return true
	}
	return false
}

func (self *Sphere) GetMaterial() Material {
	return self.Material
}
