package pathtracer

type AnimatedValue interface {
	Update(t float64)
	Get() float64
}

type AnimatedVector interface {
	Update(t float64)
	Get() *Vector3
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
