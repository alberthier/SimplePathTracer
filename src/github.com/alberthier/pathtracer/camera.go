package pathtracer

import (
	"math"
	"math/rand"
)

type Camera struct {
	position    AnimatedVector
	lookAt      AnimatedVector
	up          AnimatedVector
	vertFov     AnimatedValue
	aperture    AnimatedValue
	aspectRatio float64

	lowerLeftCorner *Vector3
	horizontal      *Vector3
	vertical        *Vector3
	lensRadius      float64
	u               *Vector3
	v               *Vector3
	w               *Vector3
}

func NewCamera(position AnimatedVector, lookAt AnimatedVector, up AnimatedVector, vertFov AnimatedValue, aspectRatio float64, aperture AnimatedValue) *Camera {
	self := &Camera{}
	self.position = position
	self.lookAt = lookAt
	self.up = up
	self.vertFov = vertFov
	self.aperture = aperture
	self.aspectRatio = aspectRatio
	return self
}

func (self *Camera) Update(t float64) {
	self.position.Update(t)
	self.lookAt.Update(t)
	self.up.Update(t)
	self.vertFov.Update(t)
	self.aperture.Update(t)

	theta := self.vertFov.Get() * math.Pi / 180.0
	focusDistance := self.position.Get().Subtract(self.lookAt.Get()).Length()
	lHeight := math.Tan(theta/2.0) * focusDistance
	lWidth := self.aspectRatio * lHeight

	self.w = self.position.Get().Subtract(self.lookAt.Get()).Unit()
	self.u = self.up.Get().Cross(self.w).Unit()
	self.v = self.w.Cross(self.u)

	self.lensRadius = self.aperture.Get() / 2.0
	self.lowerLeftCorner = self.position.Get().Subtract(self.u.Scale(lWidth)).Subtract(self.v.Scale(lHeight)).Subtract(self.w.Scale(focusDistance))
	self.horizontal = self.u.Scale(2.0 * lWidth)
	self.vertical = self.v.Scale(2.0 * lHeight)
}

func randomVectorInUnitDisk(rng *rand.Rand) *Vector3 {
	for {
		p := Vector3{2.0*rng.Float64() - 1.0, 2.0*rng.Float64() - 1.0, 0.0}
		if p.Dot(&p) < 1.0 {
			return &p
		}
	}
}

func (self *Camera) GetRay(rng *rand.Rand, s float64, t float64) *Ray {
	rnd := randomVectorInUnitDisk(rng).Scale(self.lensRadius)
	offset := self.u.Scale(rnd.X).Add(self.v.Scale(rnd.Y))
	return NewRay(self.position.Get().Add(offset),
		self.lowerLeftCorner.Add(self.horizontal.Scale(s)).Add(self.vertical.Scale(t)).Subtract(self.position.Get()).Subtract(offset))
}
