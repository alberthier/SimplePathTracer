package pathtracer

type SceneObject interface {
	HitBy(ray *Ray) bool
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

func (self *Sphere) HitBy(ray *Ray) bool {
	oc := ray.Origin.Substract(&self.Position)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * oc.Dot(ray.Direction)
	c := oc.Dot(oc) - self.Radius*self.Radius
	disc := b*b - 4.0*a*c
	return disc > 0
}
