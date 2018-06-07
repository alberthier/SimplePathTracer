package pathtracer

import "math"

type SceneObject interface {
	HitBy(ray *Ray) bool
	Color(ray *Ray) *Color
}

type ObjectBase struct {
	Position Vector3
}

type Sphere struct {
	ObjectBase
	Radius   float32
	Material *Material
}

var (
	neg1z = NewVector(0.0, 0.0, -1.0)
)

func NewSphere(x float32, y float32, z float32, radius float32, material *Material) *Sphere {
	return &Sphere{
		ObjectBase{Vector3{x, y, z}},
		radius, material}
}

func (self *Sphere) HitBy(ray *Ray) bool {
	oc := ray.Origin.Substract(&self.Position)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * oc.Dot(ray.Direction)
	c := oc.Dot(oc) - self.Radius*self.Radius
	disc := b*b - 4.0*a*c
	return disc > 0
}

func (self *Sphere) Color(ray *Ray) *Color {
	oc := ray.Origin.Substract(&self.Position)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * oc.Dot(ray.Direction)
	c := oc.Dot(oc) - self.Radius*self.Radius
	disc := b*b - 4.0*a*c
	if disc < 0 {
		return nil
	}
	t := (-b - float32(math.Sqrt(float64(disc)))) / (2.0 * a)
	norm := ray.PointAt(t).Substract(neg1z).Unit()
	return NewColor((norm.X+1.0)/2.0, (norm.Y+1.0)/2.0, (norm.Z+1.0)/2.0)
}
