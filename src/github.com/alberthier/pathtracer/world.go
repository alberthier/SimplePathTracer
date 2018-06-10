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
		Name    string  `json:"name"`
		Type    string  `json:"type"`
		Diffuse float32 `json:"diffuse"`
		Scatter float32 `json:"scatter"`
	} `json:"materials"`
	Scene struct {
		Camera struct {
			X      float32 `json:"x"`
			Y      float32 `json:"y"`
			Z      float32 `json:"z"`
			AngleX float32 `json:"ax"`
			AngleY float32 `json:"ay"`
			AngleZ float32 `json:"az"`
		} `json:"camera"`
		Objects []struct {
			Type     string  `json:"type"`
			X        float32 `json:"x"`
			Y        float32 `json:"y"`
			Z        float32 `json:"z"`
			Radius   float32 `json:"radius"`
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
		self.Materials[matData.Name] = NewMaterial(matData.Type)
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
