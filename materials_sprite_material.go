package gthree

//go:generate go run material_method_generator/main.go -materialName SpriteMaterial -materialSlug sprite_material

import (
	"github.com/gopherjs/gopherjs/js"
)

type SpriteMaterial struct {
	*js.Object
}

func NewSpriteMaterial(texture *Texture) *SpriteMaterial {
	params := NewMaterialParameters()
	params.Map = texture
	return &SpriteMaterial{
		Object: three.Get("SpriteMaterial").New(params),
	}
}
