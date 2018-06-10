package pathtracer

import (
	"math"
)

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

var vector000 = NewVector(0.0, 0.0, 0.0)
var vector111 = NewVector(1.0, 1.0, 1.0)

func NewVector(x float32, y float32, z float32) *Vector3 {
	return &Vector3{x, y, z}
}

func (self *Vector3) Add(other *Vector3) *Vector3 {
	return &Vector3{self.X + other.X, self.Y + other.Y, self.Z + other.Z}
}

func (self *Vector3) Substract(other *Vector3) *Vector3 {
	return &Vector3{self.X - other.X, self.Y - other.Y, self.Z - other.Z}
}

func (self *Vector3) Multiply(other *Vector3) *Vector3 {
	return &Vector3{self.X * other.X, self.Y * other.Y, self.Z * other.Z}
}

func (self *Vector3) Divide(other *Vector3) *Vector3 {
	return &Vector3{self.X / other.X, self.Y / other.Y, self.Z / other.Z}
}

func (self *Vector3) Scale(factor float32) *Vector3 {
	return &Vector3{self.X * factor, self.Y * factor, self.Z * factor}
}

func (self *Vector3) Dot(other *Vector3) float32 {
	return self.X*other.X + self.Y*other.Y + self.Z*other.Z
}

func (self *Vector3) Cross(other *Vector3) *Vector3 {
	return &Vector3{
		self.Y*other.Z - self.Z*other.Y,
		self.Z*other.X - self.X*other.Z,
		self.X*other.Y - self.Y*other.X}
}

func (self *Vector3) Length() float32 {
	return float32(math.Sqrt(float64(self.SquaredLength())))
}

func (self *Vector3) SquaredLength() float32 {
	return self.X*self.X + self.Y*self.Y + self.Z*self.Z
}

func (self *Vector3) Unit() *Vector3 {
	return self.Scale(1.0 / self.Length())
}
