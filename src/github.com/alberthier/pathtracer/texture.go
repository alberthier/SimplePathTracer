package pathtracer

import (
	"math"
)

type Texture interface {
	Color(point *Vector3) *Color
}

func NewTexture(typ string, color [3]float64, size float64, param1 string, param2 string, textures *map[string]Texture) Texture {
	switch typ {
	case "static":
		return NewStaticTexture(NewColor(color[0], color[1], color[2]))
	case "checker":
		texture1, ok1 := (*textures)[param1]
		texture2, ok2 := (*textures)[param2]
		if ok1 && ok2 {
			return NewCheckerTexture(size, texture1, texture2)
		}
		break
	}
	return nil
}

// Static color =======================================================

type StaticColor struct {
	color *Color
}

func NewStaticTexture(color *Color) *StaticColor {
	return &StaticColor{color}
}

func (self *StaticColor) Color(point *Vector3) *Color {
	return self.color
}

// Checker texture =======================================================

type CheckerTexture struct {
	size        float64
	evenTexture Texture
	oddTexture  Texture
}

func NewCheckerTexture(size float64, evenTexture Texture, oddTexture Texture) *CheckerTexture {
	return &CheckerTexture{size, evenTexture, oddTexture}
}

func (self *CheckerTexture) Color(point *Vector3) *Color {
	s := math.Sin(self.size*point.X) * math.Sin(self.size*point.Y) * math.Sin(self.size*point.Z)
	if s < 0 {
		return self.oddTexture.Color(point)
	} else {
		return self.evenTexture.Color(point)
	}
}
