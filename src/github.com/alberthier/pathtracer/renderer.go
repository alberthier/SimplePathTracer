package pathtracer

import (
	"image"
)

// ===================== Color

type Color struct {
	R float32
	G float32
	B float32
}

func NewColor(r float32, g float32, b float32) *Color {
	return &Color{r, g, b}
}

func (self *Color) RGBA() (r uint32, g uint32, b uint32, a uint32) {
	return uint32(self.R * 0xffff),
		uint32(self.G * 0xffff),
		uint32(self.B * 0xffff),
		0xffff
}

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

func (self *Renderer) background(ray *Ray) *Color {
	udir := ray.Direction.Unit()
	t := 0.5 * (udir.Y + 1.0)
	return NewColor((1-t)+t*0.5,
		(1-t)+t*0.7,
		(1-t)+t*1.0)
}

func (self *Renderer) Render(world *World) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, self.width, self.height))

	fwidth := float32(self.width)
	fheight := float32(self.height)
	uw := fwidth / 100.0
	uh := fheight / 100.0
	lowerLeftCorner := NewVector(-uw, -uh, -1.0)
	hSize := NewVector(2.0*uw, 0.0, 0.0)
	vSize := NewVector(0.0, 2.0*uh, 0.0)

	ray := Ray{}
	ray.Origin = NewVector(0.0, 0.0, 0.0)

	for j := self.height - 1; j >= 0; j-- {
		for i := 0; i < self.width; i++ {
			u := float32(i) / fwidth
			v := float32(j) / fheight
			ray.Direction = lowerLeftCorner.Add(hSize.Scale(u)).Add(vSize.Scale(v))
			var color *Color
			hit := false
			for _, obj := range world.Scene.Objects {
				hit = obj.HitBy(&ray)
				if hit {
					color = NewColor(1.0, 0.0, 0.0)
					break
				}
			}
			if !hit {
				color = self.background(&ray)
			}
			img.Set(i, self.height-j, color)
		}
	}

	return img
}
