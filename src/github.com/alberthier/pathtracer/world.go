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

type WorldFile struct {
	Materials []struct {
		Name    string     `json:"name"`
		Type    string     `json:"type"`
		Diffuse [3]float64 `json:"diffuse"`
		Param   float64    `json:"param"`
	} `json:"materials"`
	Scene struct {
		Camera struct {
			X        float64 `json:"x"`
			Y        float64 `json:"y"`
			Z        float64 `json:"z"`
			ToX      float64 `json:"tox"`
			ToY      float64 `json:"toy"`
			ToZ      float64 `json:"toz"`
			UpX      float64 `json:"upx"`
			UpY      float64 `json:"upy"`
			UpZ      float64 `json:"upz"`
			Fov      float64 `json:"fov"`
			Aperture float64 `json:"aperture"`
		} `json:"camera"`
		Objects []struct {
			Type     string  `json:"type"`
			X        float64 `json:"x"`
			Y        float64 `json:"y"`
			Z        float64 `json:"z"`
			Radius   float64 `json:"radius"`
			Material string  `json:"material"`
		}
	} `json:"scene"`
}

func (self *World) Load(filename string, aspectRatio float64) error {
	bytes, _ := ioutil.ReadFile(filename)
	worldFile := WorldFile{}
	err := json.Unmarshal(bytes, &worldFile)

	if err != nil {
		return errors.New("Unable to parse JSON")
	}

	camPos := NewVector(worldFile.Scene.Camera.X, worldFile.Scene.Camera.Y, worldFile.Scene.Camera.Z)
	camLookAt := NewVector(worldFile.Scene.Camera.ToX, worldFile.Scene.Camera.ToY, worldFile.Scene.Camera.ToZ)
	camUp := NewVector(worldFile.Scene.Camera.UpX, worldFile.Scene.Camera.UpY, worldFile.Scene.Camera.UpZ)

	self.Scene.Camera = NewCamera(camPos, camLookAt, camUp, worldFile.Scene.Camera.Fov, aspectRatio, worldFile.Scene.Camera.Aperture)

	self.Materials = make(map[string]Material)
	for _, matData := range worldFile.Materials {
		self.Materials[matData.Name] = NewMaterial(matData.Type, matData.Diffuse, matData.Param)
	}

	for _, objData := range worldFile.Scene.Objects {
		switch objData.Type {
		case "sphere":
			if material, ok := self.Materials[objData.Material]; ok {
				sphere := NewSphere(objData.X, objData.Y, objData.Z, objData.Radius, material)
				self.Scene.Objects = append(self.Scene.Objects, sphere)
			} else {
				fmt.Printf("Object material not found: '%s'", objData.Material)
			}
			break
		}
	}

	return nil
}
