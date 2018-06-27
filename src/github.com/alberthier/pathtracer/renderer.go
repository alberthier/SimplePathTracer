package pathtracer

import (
	"fmt"
	"image"
	"math"
	"math/rand"
	"runtime"
	"time"
)

// ===================== Color

type Color struct {
	R float64
	G float64
	B float64
}

var WhiteColor = NewColor(1.0, 1.0, 1.0)

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

type PixelColor struct {
	X     int
	Y     int
	Color *Color
}

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

func (self *Renderer) Color(rng *rand.Rand, ray *Ray, world *World, depth int) *Color {
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
			attenuation, scattered := record.object.GetMaterial().Scatter(rng, ray, &record)
			if attenuation != nil && scattered != nil {
				color := self.Color(rng, scattered, world, depth+1)
				return NewColor(attenuation.R*color.R,
					attenuation.G*color.G,
					attenuation.B*color.B)
			}
		}
		return NewColor(0.0, 0.0, 0.0)
	} else {
		return self.background(ray)
	}
}

func (self *Renderer) renderLine(channel chan *PixelColor, world *World, line int) {
	fwidth := float64(self.width)
	fheight := float64(self.height)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < self.width; i++ {
		color := NewColor(0.0, 0.0, 0.0)
		for s := 0; s < self.samplesPerPx; s++ {
			u := (float64(i) + rng.Float64()) / fwidth
			v := (float64(line) + rng.Float64()) / fheight
			ray := world.Scene.Camera.GetRay(rng, u, v)
			color.AddFrom(self.Color(rng, ray, world, 0))
		}
		color.DivideAll(float64(self.samplesPerPx))
		color.GammaCorrect()

		channel <- &PixelColor{i, self.height - line - 1, color}
	}
	channel <- &PixelColor{0, 0, nil}
}

func (self *Renderer) Render(world *World, t float64, logprefix string) image.Image {
	maxGoRoutines := runtime.NumCPU()
	goRoutinesCount := 0

	img := image.NewRGBA(image.Rect(0, 0, self.width, self.height))

	channel := make(chan *PixelColor, 100)

	world.Update(t)

	line := 0
	totalPixels := float64(self.width * self.height)
	renderedPixels := 0
	for line < self.height || goRoutinesCount != 0 {
		for goRoutinesCount < maxGoRoutines && line < self.height {
			go self.renderLine(channel, world, line)
			goRoutinesCount++
			line++
		}
		for {
			pix := <-channel
			if pix.Color != nil {
				img.Set(pix.X, pix.Y, pix.Color)
				renderedPixels++
				fmt.Printf("\r%s%.1f%%", logprefix, 100.0*float64(renderedPixels)/totalPixels)
			} else {
				goRoutinesCount--
				break
			}
		}
	}
	fmt.Printf("\r%s      \r%sOK\n", logprefix, logprefix)

	return img
}
