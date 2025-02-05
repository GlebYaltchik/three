package gthree

//go:generate go run material_method_generator/main.go -materialName MeshPhongMaterial -materialSlug mesh_phong_material

import (
	"github.com/gopherjs/gopherjs/js"
)

type MeshPhongMaterial struct {
	*js.Object

	// BumpMap     *Texture `js:"bumpMap"`
	// BumpScale   float64  `js:"bumpScale"`
	// SpecularMap *Texture `js:"specularMap"`
	// Specular    *Color   `js:"specular"`
}

func NewMeshPhongMaterial(params *MaterialParameters) *MeshPhongMaterial {
	return &MeshPhongMaterial{
		Object: three.Get("MeshPhongMaterial").New(params),
	}
}
