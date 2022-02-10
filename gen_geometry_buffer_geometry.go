package gthree

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// geometry_method_generator -geometryType BufferGeometry -geometrySlug buffer_geometry

import "github.com/gopherjs/gopherjs/js"

var _ Geometry = BufferGeometry{}

func (g BufferGeometry) ApplyMatrix4(matrix *Matrix4) {
	g.Object.Call("applyMatrix4", matrix)
}

func (g BufferGeometry) RotateX() {
	g.Object.Call("rotateX")
}

func (g BufferGeometry) RotateY() {
	g.Object.Call("rotateY")
}

func (g BufferGeometry) RotateZ() {
	g.Object.Call("rotateZ")
}

func (g BufferGeometry) Translate() {
	g.Object.Call("translate")
}

func (g BufferGeometry) Scale() {
	g.Object.Call("scale")
}

func (g BufferGeometry) LookAt() {
	g.Object.Call("lookAt")
}

func (g BufferGeometry) FromBufferGeometry(geometry Geometry) {
	g.Object.Call("fromBufferGeometry")
}

func (g BufferGeometry) Center() {
	g.Object.Call("center")
}

func (g BufferGeometry) Normalize() BufferGeometry {
	g.Object.Call("normalize")
	return g
}

func (g BufferGeometry) ComputeFaceNormals() {
	g.Object.Call("computeFaceNormals")
}

func (g BufferGeometry) ComputeVertexNormals(areaWeighted bool) {
	g.Object.Call("computeVertexNormals", areaWeighted)
}

func (g BufferGeometry) ComputeFlatVertexNormals() {
	g.Object.Call("computeFlatVertexNormals")
}

func (g BufferGeometry) ComputeMorphNormals() {
	g.Object.Call("computeMorphNormals")
}

func (g BufferGeometry) ComputeLineDistances() {
	g.Object.Call("computeLineDistances")
}

func (g BufferGeometry) ComputeBoundingBox() {
	g.Object.Call("computeBoundingBox")
}

func (g BufferGeometry) ComputeBoundingSphere() {
	g.Object.Call("computeBoundingSphere")
}

func (g BufferGeometry) Merge(geometry Geometry, matrix Matrix4, materialIndexOffset float64) {
	g.Object.Call("merge", geometry, matrix, materialIndexOffset)
}

func (g BufferGeometry) MergeMesh(mesh Mesh) {
	g.Object.Call("mergeMesh", mesh.getInternalObject())
}

func (g BufferGeometry) MergeVertices() {
	g.Object.Call("mergeVertices")
}

func (g BufferGeometry) SortFacesByMaterialIndex() {
	g.Object.Call("sortFacesByMaterialIndex")
}

func (g BufferGeometry) ToJSON() interface{} {
	return g.Object.Call("toJSON")
}

// func (g BufferGeometry) Clone() BufferGeometry {
// 	return g.Object.Call("clone")
// }

func (g BufferGeometry) Copy(source Object3D, recursive bool) *BufferGeometry {
	return &BufferGeometry{Object: g.getInternalObject().Call("copy", source.getInternalObject(), recursive)}
}

func (g BufferGeometry) Dispose() {
	g.Object.Call("dispose")
}

func (g BufferGeometry) getInternalObject() *js.Object {
	return g.Object
}
