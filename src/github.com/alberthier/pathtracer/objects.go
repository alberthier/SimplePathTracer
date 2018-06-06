package pathtracer

type SceneObject interface {
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
