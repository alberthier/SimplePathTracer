package pathtracer

type Camera struct {
	Position Vector3
	AngleX   float64
	AngleY   float64
	AngleZ   float64
}

func NewCamera(x float64, y float64, z float64, angleX float64, angleY float64, angleZ float64) *Camera {
	return &Camera{Vector3{x, y, z}, angleX, angleY, angleZ}
}
