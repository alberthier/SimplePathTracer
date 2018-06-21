package pathtracer

import (
	"math"
)

type AnimatedValue interface {
	Update(t float64)
	Get() float64
	Clone() AnimatedValue
}

type AnimatedVector interface {
	Update(t float64)
	Get() *Vector3
	Clone() AnimatedVector
}

// FixedValue ===========================================

type FixedValue struct {
	value float64
}

func NewFixedValue(value float64) *FixedValue {
	return &FixedValue{value}
}

func (self *FixedValue) Update(t float64) {
}

func (self *FixedValue) Get() float64 {
	return self.value
}

func (self *FixedValue) Clone() AnimatedValue {
	return NewFixedValue(self.value)
}

// SinValue ===========================================

type SinValue struct {
	value float64
	scale float64
	speed float64
}

func NewSinValue(scale float64, speed float64) *SinValue {
	return &SinValue{0.0, scale, speed}
}

func (self *SinValue) Update(t float64) {
	self.value = self.scale * math.Sin(self.speed*t/100.0)
}

func (self *SinValue) Get() float64 {
	return self.value
}

func (self *SinValue) Clone() AnimatedValue {
	return NewSinValue(self.scale, self.speed)
}

// FixedVector3 ===========================================

type FixedVector3 struct {
	value Vector3
}

func NewFixedVector3(x float64, y float64, z float64) *FixedVector3 {
	return &FixedVector3{Vector3{x, y, z}}
}

func (self *FixedVector3) Update(t float64) {
}

func (self *FixedVector3) Get() *Vector3 {
	return &self.value
}

func (self *FixedVector3) Clone() AnimatedVector {
	return NewFixedVector3(self.value.X, self.value.Y, self.value.Z)
}

// CircularPositionVector3 ===================================

type CircularYPositionVector3 struct {
	value  Vector3
	center Vector3
	radius float64
	speed  float64
}

func NewCircularYPositionVector3(cx float64, cy float64, cz float64, radius float64, speed float64) *CircularYPositionVector3 {
	return &CircularYPositionVector3{Vector3{}, Vector3{cx, cy, cz}, radius, speed}
}

func (self *CircularYPositionVector3) Update(t float64) {
	self.value.X = self.center.X + math.Cos(self.speed*t/100.0)*self.radius
	self.value.Y = self.center.Y
	self.value.Z = self.center.Z + math.Sin(self.speed*t/100.0)*self.radius
}

func (self *CircularYPositionVector3) Get() *Vector3 {
	return &self.value
}

func (self *CircularYPositionVector3) Clone() AnimatedVector {
	return NewCircularYPositionVector3(self.center.X, self.center.Y, self.center.Z, self.radius, self.speed)
}
