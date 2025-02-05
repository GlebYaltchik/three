package gthree

//go:generate go run object3d_method_generator/main.go -typeName Points -typeSlug points

import (
	"github.com/gopherjs/gopherjs/js"
)

// A class for displaying points. The points are rendered by the WebGLRenderer using gl.POINTS.
type Points struct {
	*js.Object

	// ID               int             `js:"id"`
	// Position         *Vector3        `js:"position"`
	Geometry *BufferGeometry `js:"geometry"`
	Material Material        `js:"material"`
}

// NewLine creates a new material. If Material is nil, three.js will assign a randomized material to the line o_O.
func NewPoints(geom Geometry, material Material) *Points {
	return &Points{
		Object: three.Get("Points").New(geom, material),
	}
}

// Get intersections between a casted ray and this Points. Raycaster.intersectObject will call this method.
func (p *Points) Raycast(rc Raycaster, intersects *js.Object) {
	p.Object.Call("raycast", rc, intersects)
}
