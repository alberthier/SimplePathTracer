package pathtracer

import (
	"image"
	"math"
	"math/rand"
)

// ===================== Color

type Color struct {
	R float64
	G float64
	B float64
}

func NewColor(r float64, g float64, b float64) *Color {
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

func (self *Color) DivideAll(val float64) {
	self.R /= val
	self.G /= val
	self.B /= val
}

func (self *Color) GammaCorrect() {
	self.R = math.Sqrt(self.R)
	self.G = math.Sqrt(self.G)
	self.B = math.Sqrt(self.B)
}

// ===================== Ray

type Ray struct {
	Origin    *Vector3
	Direction *Vector3
}

func NewRay(origin *Vector3, direction *Vector3) *Ray {
	return &Ray{origin, direction}
}

func (self *Ray) PointAt(t float64) *Vector3 {
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
	tmax := math.MaxFloat64
	hitSomething := false
	for _, obj := range world.Scene.Objects {
		hit := obj.HitBy(ray, 0.001, tmax, &record)
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

	fwidth := float64(self.width)
	fheight := float64(self.height)

	for j := self.height - 1; j >= 0; j-- {
		for i := 0; i < self.width; i++ {
			color := NewColor(0.0, 0.0, 0.0)
			for s := 0; s < self.samplesPerPx; s++ {
				u := (float64(i) + rand.Float64()) / fwidth
				v := (float64(j) + rand.Float64()) / fheight
				ray := world.Scene.Camera.GetRay(u, v)
				color.AddFrom(self.Color(ray, world, 0))
			}
			color.DivideAll(float64(self.samplesPerPx))
			color.GammaCorrect()
			img.Set(i, self.height-j-1, color)
		}
	}

	return img
}
