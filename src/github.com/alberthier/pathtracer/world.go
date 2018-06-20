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
			X      float64 `json:"x"`
			Y      float64 `json:"y"`
			Z      float64 `json:"z"`
			AngleX float64 `json:"ax"`
			AngleY float64 `json:"ay"`
			AngleZ float64 `json:"az"`
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

func (self *World) Load(filename string) error {
	bytes, _ := ioutil.ReadFile(filename)
	worldFile := WorldFile{}
	err := json.Unmarshal(bytes, &worldFile)

	if err != nil {
		return errors.New("Unable to parse JSON")
	}

	self.Scene.Camera = NewCamera(
		worldFile.Scene.Camera.X,
		worldFile.Scene.Camera.Y,
		worldFile.Scene.Camera.Z,
		worldFile.Scene.Camera.AngleX,
		worldFile.Scene.Camera.AngleY,
		worldFile.Scene.Camera.AngleZ)

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
