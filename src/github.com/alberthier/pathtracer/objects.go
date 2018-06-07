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
	Color(record *HitRecord) *Color
}

type ObjectBase struct {
	Position Vector3
}

type Sphere struct {
	ObjectBase
	Radius   float32
	Material *Material
}

func NewSphere(x float32, y float32, z float32, radius float32, material *Material) *Sphere {
	return &Sphere{
		ObjectBase{Vector3{x, y, z}},
		radius, material}
}

func (self *Sphere) HitBy(ray *Ray, tmin float32, tmax float32, record *HitRecord) bool {
	oc := ray.Origin.Substract(&self.Position)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * oc.Dot(ray.Direction)
	c := oc.Dot(oc) - self.Radius*self.Radius
	disc := b*b - 4.0*a*c
	if disc >= 0 {
		sd := float32(math.Sqrt(float64(disc)))
		da := 2 * a
		t := (-b - sd) / da
		if t <= tmin || tmax <= t {
			t := (-b + sd) / da
			if t <= tmin || tmax <= t {
				return false
			}
		}
		record.t = t
		record.point = ray.PointAt(t)
		record.normal = record.point.Substract(&self.Position).Unit()
		record.object = self
		return true
	}
	return false
}

func (self *Sphere) Color(record *HitRecord) *Color {
	return NewColor((record.normal.X+1.0)/2.0, (record.normal.Y+1.0)/2.0, (record.normal.Z+1.0)/2.0)
}
