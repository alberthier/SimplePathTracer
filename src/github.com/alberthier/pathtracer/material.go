package pathtracer

type Material struct {
	Type      int        `json:"type"`
	Albedo    [3]float32 `json:"albedo"`
	Emission  [3]float32 `json:"emission"`
	Roughness float32    `json:"roughness"`
	Ri        float32    `json:"ri"`
}

func NewMaterial(tp string) *Material {
	switch tp {
	case "lambert":
		return &Material{}
	case "metal":
		return &Material{}
	case "dielectric":
		return &Material{}
	}
	return nil
}

type LambertMat struct {
	Material
}
