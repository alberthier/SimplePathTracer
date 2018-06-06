package pathtracer

import (
	"image"
)

// ===================== Ray

type Ray struct {
	Origin    *Vector3
	Direction *Vector3
}

func NewRay(origin *Vector3, direction *Vector3) *Ray {
	return &Ray{origin, direction}
}

func (self *Ray) PointAt(t float32) *Vector3 {
	return self.Origin.Add(self.Direction.Scale(t))
}

// ===================== Renderer

type Renderer struct {
	width  int
	height int
}

func NewRenderer(width int, height int) *Renderer {
	return &Renderer{width, height}
}

func (self *Renderer) Render(world *World) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, self.width, self.height))

	return img
}
