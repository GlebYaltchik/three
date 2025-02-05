package gthree

//go:generate go run geometry_method_generator/main.go -geometryType SphereGeometry -geometrySlug sphere_geometry

import (
	"math"

	"github.com/gopherjs/gopherjs/js"
)

type SphereGeometry struct {
	*js.Object

	Radius float64 `js:"radius"`
	// Segments    int     `js:"segments"` // Not sure if available.
	ThetaStart  float64 `js:"thetaStart"`
	ThetaLength float64 `js:"thetaLength"`
}

type SphereGeometryParameters struct {
	Radius         float64
	WidthSegments  int
	HeightSegments int
	ThetaStart     float64
	ThetaLength    float64
	PhiStart       float64
	PhiLength      float64
}

// NewSphereGeometry creates a new BoxGeometry.
func NewSphereGeometry(params SphereGeometryParameters) SphereGeometry {
	if params.ThetaLength == 0 {
		params.ThetaLength = math.Pi
	}
	if params.PhiLength == 0 {
		params.PhiLength = 2 * math.Pi
	}
	if params.WidthSegments == 0 {
		params.WidthSegments = 8
	}
	if params.HeightSegments == 0 {
		params.HeightSegments = 6
	}
	if params.Radius == 0 {
		params.Radius = 1
	}
	return SphereGeometry{
		Object: three.Get("SphereGeometry").New(
			params.Radius,
			params.WidthSegments,
			params.HeightSegments,
			params.PhiStart,
			params.PhiLength,
			params.ThetaStart,
			params.ThetaLength,
		),
	}
}
