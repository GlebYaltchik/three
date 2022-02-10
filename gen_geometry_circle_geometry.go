package gthree

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// geometry_method_generator -geometryType CircleGeometry -geometrySlug circle_geometry

import "github.com/gopherjs/gopherjs/js"

var _ Geometry = CircleGeometry{}

func (g CircleGeometry) ApplyMatrix4(matrix *Matrix4) {
	g.Object.Call("applyMatrix4", matrix)
}

func (g CircleGeometry) RotateX() {
	g.Object.Call("rotateX")
}

func (g CircleGeometry) RotateY() {
	g.Object.Call("rotateY")
}

func (g CircleGeometry) RotateZ() {
	g.Object.Call("rotateZ")
}

func (g CircleGeometry) Translate() {
	g.Object.Call("translate")
}

func (g CircleGeometry) Scale() {
	g.Object.Call("scale")
}

func (g CircleGeometry) LookAt() {
	g.Object.Call("lookAt")
}

func (g CircleGeometry) FromBufferGeometry(geometry Geometry) {
	g.Object.Call("fromBufferGeometry")
}

func (g CircleGeometry) Center() {
	g.Object.Call("center")
}

func (g CircleGeometry) Normalize() CircleGeometry {
	g.Object.Call("normalize")
	return g
}

func (g CircleGeometry) ComputeFaceNormals() {
	g.Object.Call("computeFaceNormals")
}

func (g CircleGeometry) ComputeVertexNormals(areaWeighted bool) {
	g.Object.Call("computeVertexNormals", areaWeighted)
}

func (g CircleGeometry) ComputeFlatVertexNormals() {
	g.Object.Call("computeFlatVertexNormals")
}

func (g CircleGeometry) ComputeMorphNormals() {
	g.Object.Call("computeMorphNormals")
}

func (g CircleGeometry) ComputeLineDistances() {
	g.Object.Call("computeLineDistances")
}

func (g CircleGeometry) ComputeBoundingBox() {
	g.Object.Call("computeBoundingBox")
}

func (g CircleGeometry) ComputeBoundingSphere() {
	g.Object.Call("computeBoundingSphere")
}

func (g CircleGeometry) Merge(geometry Geometry, matrix Matrix4, materialIndexOffset float64) {
	g.Object.Call("merge", geometry, matrix, materialIndexOffset)
}

func (g CircleGeometry) MergeMesh(mesh Mesh) {
	g.Object.Call("mergeMesh", mesh.getInternalObject())
}

func (g CircleGeometry) MergeVertices() {
	g.Object.Call("mergeVertices")
}

func (g CircleGeometry) SortFacesByMaterialIndex() {
	g.Object.Call("sortFacesByMaterialIndex")
}

func (g CircleGeometry) ToJSON() interface{} {
	return g.Object.Call("toJSON")
}

// func (g CircleGeometry) Clone() CircleGeometry {
// 	return g.Object.Call("clone")
// }

func (g CircleGeometry) Copy(source Object3D, recursive bool) *CircleGeometry {
	return &CircleGeometry{Object: g.getInternalObject().Call("copy", source.getInternalObject(), recursive)}
}

func (g CircleGeometry) Dispose() {
	g.Object.Call("dispose")
}

func (g CircleGeometry) getInternalObject() *js.Object {
	return g.Object
}
