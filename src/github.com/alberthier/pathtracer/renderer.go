package pathtracer

import (
	"image"
	"math"
	"math/rand"
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

func (self *Color) AddFrom(other *Color) {
	self.R += other.R
	self.G += other.G
	self.B += other.B
}

func (self *Color) DivideAll(val float32) {
	self.R /= val
	self.G /= val
	self.B /= val
}

func (self *Color) GammaCorrect() {
	self.R = float32(math.Sqrt(float64(self.R)))
	self.G = float32(math.Sqrt(float64(self.G)))
	self.B = float32(math.Sqrt(float64(self.B)))
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
	width        int
	height       int
	samplesPerPx int
}

func NewRenderer(width int, height int, samplesPerPx int) *Renderer {
	return &Renderer{width, height, samplesPerPx}
}

func (self *Renderer) background(ray *Ray) *Color {
	udir := ray.Direction.Unit()
	t := 0.5 * (udir.Y + 1.0)
	return NewColor((1-t)+t*0.5,
		(1-t)+t*0.7,
		(1-t)+t*1.0)
}

func (self *Renderer) Color(ray *Ray, world *World, depth int) *Color {
	record := HitRecord{}
	tmax := float32(math.MaxFloat32)
	hitSomething := false
	for _, obj := range world.Scene.Objects {
		hit := obj.HitBy(ray, float32(0.001), tmax, &record)
		if hit {
			hitSomething = true
			tmax = record.t
		}
	}
	if hitSomething {
		if depth < 50 {
			attenuation, scattered := record.object.GetMaterial().Scatter(ray, &record)
			if attenuation != nil && scattered != nil {
				color := self.Color(scattered, world, depth+1)
				return NewColor(attenuation.X*color.R,
					attenuation.Y*color.G,
					attenuation.Z*color.B)
			}
		}
		return NewColor(0.0, 0.0, 0.0)
	} else {
		return self.background(ray)
	}
}

func (self *Renderer) Render(world *World) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, self.width, self.height))

	fwidth := float32(self.width)
	fheight := float32(self.height)
	lowerLeftCorner := NewVector(-2.0, -1.0, -1.0)
	hSize := NewVector(4.0, 0.0, 0.0)
	vSize := NewVector(0.0, 2.0, 0.0)

	ray := Ray{}
	ray.Origin = NewVector(0.0, 0.0, 0.0)

	for j := self.height - 1; j >= 0; j-- {
		for i := 0; i < self.width; i++ {
			color := NewColor(0.0, 0.0, 0.0)
			for s := 0; s < self.samplesPerPx; s++ {
				u := (float32(i) + rand.Float32()) / fwidth
				v := (float32(j) + rand.Float32()) / fheight
				ray.Direction = lowerLeftCorner.Add(hSize.Scale(u)).Add(vSize.Scale(v)).Substract(ray.Origin)
				color.AddFrom(self.Color(&ray, world, 0))
			}
			color.DivideAll(float32(self.samplesPerPx))
			color.GammaCorrect()
			img.Set(i, self.height-j-1, color)
		}
	}

	return img
}
