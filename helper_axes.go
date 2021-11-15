package three

import "github.com/gopherjs/gopherjs/js"

type AxesHelper struct {
	*js.Object

	// Copied from LineSegments struct:

	ID               int             `js:"id"`
	Position         *Vector3        `js:"position"`
	Rotation         *Euler          `js:"rotation"`
	Geometry         *BufferGeometry `js:"geometry"`
	Material         Material        `js:"material"`
	MatrixAutoUpdate bool            `js:"matrixAutoUpdate"`
	RenderOrder      int             `js:"renderOrder"`
	Visible          bool            `js:"visible"`
}

func NewAxesHelper(size float64) AxesHelper {
	if size == 0 {
		size = 1
	}
	return AxesHelper{
		Object: three.Get("AxesHelper").New(size),
	}
}

func (g AxesHelper) SetColors(xaxis, yaxis, zaxis Color) AxesHelper {
	g.Object.Call("setColor", xaxis, yaxis, zaxis)
	return g
}

// Disposes of the internally-created material and geometry used by this helper.
func (g AxesHelper) Dispose() {
	g.Object.Call("dispose")
}
