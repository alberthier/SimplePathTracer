package pathtracer

import (
	"math"
	"math/rand"
)

type Camera struct {
	position        Vector3
	lowerLeftCorner Vector3
	horizontal      Vector3
	vertical        Vector3
	lensRadius      float64
	u               Vector3
	v               Vector3
	w               Vector3
}

func NewCamera(position *Vector3, lookAt *Vector3, up *Vector3, vertFov float64, aspectRatio float64, aperture float64) *Camera {
	theta := vertFov * math.Pi / 180.0
	focusDistance := position.Subtract(lookAt).Length()
	lHeight := math.Tan(theta/2.0) * focusDistance
	lWidth := aspectRatio * lHeight

	self := &Camera{}

	self.w = *position.Subtract(lookAt).Unit()
	self.u = *up.Cross(&self.w).Unit()
	self.v = *self.w.Cross(&self.u)

	self.lensRadius = aperture / 2.0
	self.position = *position
	self.lowerLeftCorner = *self.position.Subtract(self.u.Scale(lWidth)).Subtract(self.v.Scale(lHeight)).Subtract(self.w.Scale(focusDistance))
	self.horizontal = *self.u.Scale(2.0 * lWidth)
	self.vertical = *self.v.Scale(2.0 * lHeight)

	return self
}

func randomVectorInUnitDisk() *Vector3 {
	for {
		p := Vector3{2.0*rand.Float64() - 1.0, 2.0*rand.Float64() - 1.0, 0.0}
		if p.Dot(&p) < 1.0 {
			return &p
		}
	}
}

func (self *Camera) GetRay(s float64, t float64) *Ray {
	rnd := randomVectorInUnitDisk().Scale(self.lensRadius)
	offset := self.u.Scale(rnd.X).Add(self.v.Scale(rnd.Y))
	return NewRay(self.position.Add(offset),
		self.lowerLeftCorner.Add(self.horizontal.Scale(s)).Add(self.vertical.Scale(t)).Subtract(&self.position).Subtract(offset))
}
