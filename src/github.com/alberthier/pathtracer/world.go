package pathtracer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type World struct {
	Materials map[string]Material
	Scene     struct {
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
	Materials []struct {
		Name    string     `json:"name"`
		Type    string     `json:"type"`
		Diffuse [3]float64 `json:"diffuse"`
		Param   float64    `json:"param"`
	} `json:"materials"`
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

func newAnimatedValue(v *FileValue) AnimatedValue {
	if len(v.Anim) != 0 {
	}
	return NewFixedValue(v.Value)
}

func newAnimatedVector(v *FileVector) AnimatedVector {
	if len(v.Anim) != 0 {
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

	camPos := newAnimatedVector(&worldFile.Scene.Camera.Position)
	camLookAt := newAnimatedVector(&worldFile.Scene.Camera.LookAt)
	camUp := newAnimatedVector(&worldFile.Scene.Camera.Up)
	camFov := newAnimatedValue(&worldFile.Scene.Camera.Fov)
	camAperture := newAnimatedValue(&worldFile.Scene.Camera.Aperture)
	self.Scene.Camera = NewCamera(camPos, camLookAt, camUp, camFov, aspectRatio, camAperture)

	self.Materials = make(map[string]Material)
	for _, matData := range worldFile.Materials {
		self.Materials[matData.Name] = NewMaterial(matData.Type, matData.Diffuse, matData.Param)
	}

	for _, objData := range worldFile.Scene.Objects {
		switch objData.Type {
		case "sphere":
			if material, ok := self.Materials[objData.Material]; ok {
				pos := newAnimatedVector(&objData.Position)
				radius := newAnimatedValue(&objData.Radius)
				sphere := NewSphere(pos, radius, material)
				self.Scene.Objects = append(self.Scene.Objects, sphere)
			} else {
				fmt.Printf("Object material not found: '%s'", objData.Material)
			}
			break
		}
	}

	return nil
}
