package pathtracer

import (
	"math"
)

type Camera struct {
	position        Vector3
	lowerLeftCorner Vector3
	horizontal      Vector3
	vertical        Vector3
}

func NewCamera(position *Vector3, lookAt *Vector3, up *Vector3, vertFov float64, aspectRatio float64) *Camera {
	theta := vertFov * math.Pi / 180.0
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspectRatio * halfHeight

	w := position.Substract(lookAt).Unit()
	u := up.Cross(w).Unit()
	v := w.Cross(u)

	self := &Camera{}
	self.position = *position
	self.lowerLeftCorner = *self.position.Substract(u.Scale(halfWidth)).Substract(v.Scale(halfHeight)).Substract(w)
	self.horizontal = *u.Scale(2.0 * halfWidth)
	self.vertical = *v.Scale(2.0 * halfHeight)

	return self
}

func (self *Camera) GetRay(u float64, v float64) *Ray {
	return NewRay(&self.position,
		self.lowerLeftCorner.Add(self.horizontal.Scale(u)).Add(self.vertical.Scale(v)).Substract(&self.position))
}
