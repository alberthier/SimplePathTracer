package pathtracer

import (
	"math"
)

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

var NullVector = NewVector(0.0, 0.0, 0.0)
var UnitVector = NewVector(1.0, 1.0, 1.0)

func NewVector(x float64, y float64, z float64) *Vector3 {
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

func (self *Vector3) Scale(factor float64) *Vector3 {
	return &Vector3{self.X * factor, self.Y * factor, self.Z * factor}
}

func (self *Vector3) Dot(other *Vector3) float64 {
	return self.X*other.X + self.Y*other.Y + self.Z*other.Z
}

func (self *Vector3) Cross(other *Vector3) *Vector3 {
	return &Vector3{
		self.Y*other.Z - self.Z*other.Y,
		self.Z*other.X - self.X*other.Z,
		self.X*other.Y - self.Y*other.X}
}

func (self *Vector3) Length() float64 {
	return math.Sqrt(self.SquaredLength())
}

func (self *Vector3) SquaredLength() float64 {
	return self.X*self.X + self.Y*self.Y + self.Z*self.Z
}

func (self *Vector3) Unit() *Vector3 {
	return self.Scale(1.0 / self.Length())
}
