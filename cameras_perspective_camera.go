package three

//go:generate go run camera_method_generator/main.go -typeName PerspectiveCamera -typeSlug perspective_camera

import "github.com/gopherjs/gopherjs/js"

type PerspectiveCameraPosition struct {
	Z int
}

type PerspectiveCamera struct {
	*js.Object
	// Use Set method to set up vector.
	Up               Vector3 `js:"up"`
	Position         Vector3 `js:"position"`
	MatrixAutoUpdate bool    `js:"matrixAutoUpdate"`
	Aspect           float64 `js:"aspect"`
}

// Assert PerspectiveCamera implements Camera.
var _ Camera = PerspectiveCamera{}

func NewPerspectiveCamera(fov, aspect, near, far float64) PerspectiveCamera {
	return PerspectiveCamera{Object: three.Get("PerspectiveCamera").New(fov, aspect, near, far)}
}

func (c PerspectiveCamera) SetFocalLength(focalLength float64) {
	c.Object.Call("setFocalLength", focalLength)
}

func (c PerspectiveCamera) GetFocalLength() float64 {
	return c.Object.Call("getFocalLength").Float()
}

func (c PerspectiveCamera) GetEffectiveFOV() float64 {
	return c.Object.Call("getEffectiveFOV").Float()
}

func (c PerspectiveCamera) GetFilmWidth() float64 {
	return c.Object.Call("getFilmWidth").Float()
}

func (c PerspectiveCamera) GetFilmHeight() float64 {
	return c.Object.Call("getFilmHeight").Float()
}

func (c PerspectiveCamera) SetViewOffset(fullWidth, fullHeight, x, y, width, height float64) {
	c.Object.Call("setViewOffset", fullWidth, fullHeight, x, y, width, height)
}

// SetUp sets the up direction for the camera.
//
// It is the equivalent to c.Up.Set(v.X, v.Y, v.Z)
func (c PerspectiveCamera) SetUp(v Vector3) {
	c.Up.Set(v.X, v.Y, v.Z)
}

func (c PerspectiveCamera) ClearViewOffset() {
	c.Object.Call("clearViewOffset")
}

func (c PerspectiveCamera) UpdateProjectionMatrix() {
	c.Object.Call("updateProjectionMatrix")
}

func (c PerspectiveCamera) ToJSON(meta interface{}) interface{} {
	return c.Object.Call("toJSON", meta).Interface()
}

func (c PerspectiveCamera) LookAt(x, y, z float64) {
	c.Object.Call("lookAt", x, y, z)
}
