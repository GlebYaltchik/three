package three

import "github.com/gopherjs/gopherjs/js"

// TextureLoader loads a texture from file.
// See https://threejs.org/docs/#api/en/loaders/TextureLoader
type TextureLoader struct {
	*js.Object

	CrossOrigin string `js:"crossOrigin"`
}

// NewTextureLoader creates a new texture loader.
func NewTextureLoader() *TextureLoader {
	return &TextureLoader{
		Object: three.Get("TextureLoader").New(),
	}
}

// Load loads texture from url image.
func (t *TextureLoader) Load(url string, fn func(*js.Object)) Texture {
	return Texture{
		Object: t.Call("load", url, fn),
	}
}
