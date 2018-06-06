package pathtracer

type Camera struct {
	Position Vector3
	AngleX   float32
	AngleY   float32
	AngleZ   float32
}

func NewCamera(x float32, y float32, z float32, angleX float32, angleY float32, angleZ float32) *Camera {
	return &Camera{Vector3{x, y, z}, angleX, angleY, angleZ}
}
