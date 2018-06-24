package pathtracer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type World struct {
	Textures         map[string]Texture
	Materials        map[string]Material
	ValueAnimations  map[string]AnimatedValue
	VectorAnimations map[string]AnimatedVector
	Scene            struct {
		Camera  *Camera
		Objects []SceneObject
	}
}

func NewWorld() *World {
	return &World{}
}

type FileValue struct {
	Value float64 `json:"value"`
	Anim  string  `json:"anim"`
}

type FileVector struct {
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Z    float64 `json:"z"`
	Anim string  `json:"anim"`
}

type WorldFile struct {
	Textures []struct {
		Type     string     `json:"type"`
		Name     string     `json:"name"`
		Color    [3]float64 `json:"color"`
		Size     float64    `json:"size"`
		Texture1 string     `json:"texture1"`
		Texture2 string     `json:"texture2"`
	} `json:"textures"`
	Materials []struct {
		Name    string  `json:"name"`
		Type    string  `json:"type"`
		Texture string  `json:"texture"`
		Param   float64 `json:"param"`
	} `json:"materials"`
	Animations []struct {
		Name   string  `json:"name"`
		Type   string  `json:"type"`
		Cx     float64 `json:"cx"`
		Cy     float64 `json:"cy"`
		Cz     float64 `json:"cz"`
		Radius float64 `json:"radius"`
		Speed  float64 `json:"speed"`
		Scale  float64 `json:"scale"`
	} `json:"animations"`
	Scene struct {
		Camera struct {
			Position FileVector `json:"position"`
			LookAt   FileVector `json:"lookat"`
			Up       FileVector `json:"up"`
			Fov      FileValue  `json:"fov"`
			Aperture FileValue  `json:"aperture"`
		} `json:"camera"`
		Objects []struct {
			Type     string     `json:"type"`
			Position FileVector `json:"position"`
			Radius   FileValue  `json:"radius"`
			Material string     `json:"material"`
		}
	} `json:"scene"`
}

func (self *World) newAnimatedValue(v *FileValue) AnimatedValue {
	if len(v.Anim) != 0 {
		if anim, ok := self.ValueAnimations[v.Anim]; ok {
			return anim.Clone()
		}
	}
	return NewFixedValue(v.Value)
}

func (self *World) newAnimatedVector(v *FileVector) AnimatedVector {
	if len(v.Anim) != 0 {
		if anim, ok := self.VectorAnimations[v.Anim]; ok {
			return anim.Clone()
		}
	}
	return NewFixedVector3(v.X, v.Y, v.Z)
}

func (self *World) Load(filename string, aspectRatio float64) error {
	bytes, _ := ioutil.ReadFile(filename)
	worldFile := WorldFile{}
	err := json.Unmarshal(bytes, &worldFile)

	if err != nil {
		return errors.New("Unable to parse JSON")
	}

	self.Textures = make(map[string]Texture)
	for _, texData := range worldFile.Textures {
		self.Textures[texData.Name] = NewTexture(texData.Type, texData.Color, texData.Size, texData.Texture1, texData.Texture2, &self.Textures)
	}
	self.Materials = make(map[string]Material)
	for _, matData := range worldFile.Materials {
		texture := self.Textures[matData.Texture]
		self.Materials[matData.Name] = NewMaterial(matData.Type, texture, matData.Param)
	}
	self.VectorAnimations = make(map[string]AnimatedVector)
	self.ValueAnimations = make(map[string]AnimatedValue)
	for _, animData := range worldFile.Animations {
		switch animData.Type {
		case "circularPosition":
			self.VectorAnimations[animData.Name] = NewCircularYPositionVector3(animData.Cx, animData.Cy, animData.Cz, animData.Radius, animData.Speed)
			break
		case "sinValue":
			self.ValueAnimations[animData.Name] = NewSinValue(animData.Scale, animData.Speed)
			break
		}
	}
	camPos := self.newAnimatedVector(&worldFile.Scene.Camera.Position)
	camLookAt := self.newAnimatedVector(&worldFile.Scene.Camera.LookAt)
	camUp := self.newAnimatedVector(&worldFile.Scene.Camera.Up)
	camFov := self.newAnimatedValue(&worldFile.Scene.Camera.Fov)
	camAperture := self.newAnimatedValue(&worldFile.Scene.Camera.Aperture)
	self.Scene.Camera = NewCamera(camPos, camLookAt, camUp, camFov, aspectRatio, camAperture)
	for _, objData := range worldFile.Scene.Objects {
		switch objData.Type {
		case "sphere":
			if material, ok := self.Materials[objData.Material]; ok {
				pos := self.newAnimatedVector(&objData.Position)
				radius := self.newAnimatedValue(&objData.Radius)
				sphere := NewSphere(pos, radius, material)
				self.Scene.Objects = append(self.Scene.Objects, sphere)
			} else {
				fmt.Printf("Object material not found: '%s'\n", objData.Material)
			}
			break
		}
	}

	return nil
}

func (self *World) Update(t float64) {
	self.Scene.Camera.Update(t)
	for _, obj := range self.Scene.Objects {
		obj.Update(t)
	}
}
