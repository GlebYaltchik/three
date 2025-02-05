package gthree

//go:generate go run object3d_method_generator/main.go -typeName TextSprite -typeSlug text_sprite

import (
	"github.com/gopherjs/gopherjs/js"
)

type TextSprite struct {
	*js.Object

	Text string

	ID               int             `js:"id"`
	Position         *Vector3        `js:"position"`
	Rotation         *Euler          `js:"rotation"`
	Geometry         *BufferGeometry `js:"geometry"`
	MatrixAutoUpdate bool            `js:"matrixAutoUpdate"`
	RenderOrder      int             `js:"renderOrder"`
	Visible          bool            `js:"visible"`
}

var _ Object3D = &TextSprite{}

type TextSpriteParameters struct {
	*js.Object
	TextSize       int                 `js:"textSize"`
	RedrawInterval int                 `js:"redrawInterval"`
	Material       *TextSpriteMaterial `js:"material"`
	Texture        *TextTextureParams  `js:"texture"`
}

type TextSpriteMaterial struct {
	*js.Object
	Color *Color `js:"color"`
}

type TextTextureParams struct {
	*js.Object
	Text       string `js:"text"`
	FontFamily string `js:"fontFamily"`
	Align      string `js:"align"`
}

func NewTextSprite(text, align string, textSize int, color *Color) *TextSprite {
	material := &TextSpriteMaterial{
		Object: js.Global.Get("THREE").Get("SpriteMaterial").New(),
	}
	material.Color = color

	textureParams := &TextTextureParams{
		Object: js.Global.Get("THREE").Get("TextTexture").New(),
	}
	textureParams.Text = text
	textureParams.Align = align

	params := &TextSpriteParameters{
		Object: js.Global.Get("Object").New(),
	}
	params.Object.Set("texture", textureParams)
	params.Object.Set("material", material)
	params.Object.Set("textSize", textSize)

	return &TextSprite{
		Object: js.Global.Get("THREE").Get("TextSprite").New(params),
	}
}
