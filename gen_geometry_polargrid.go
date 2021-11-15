package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// geometry_method_generator -geometryType PolarGridHelper -geometrySlug polargrid

import "github.com/gopherjs/gopherjs/js"

func (g PolarGridHelper) ApplyMatrix(matrix *Matrix4) {
	g.Object.Call("applyMatrix", matrix)
}

func (g PolarGridHelper) RotateX() {
	g.Object.Call("rotateX")
}

func (g PolarGridHelper) RotateY() {
	g.Object.Call("rotateY")
}

func (g PolarGridHelper) RotateZ() {
	g.Object.Call("rotateZ")
}

func (g PolarGridHelper) Translate() {
	g.Object.Call("translate")
}

func (g PolarGridHelper) Scale() {
	g.Object.Call("scale")
}

func (g PolarGridHelper) LookAt() {
	g.Object.Call("lookAt")
}

func (g PolarGridHelper) FromBufferGeometry(geometry Geometry) {
	g.Object.Call("fromBufferGeometry")
}

func (g PolarGridHelper) Center() {
	g.Object.Call("center")
}

func (g PolarGridHelper) Normalize() PolarGridHelper {
	g.Object.Call("normalize")
	return g
}

func (g PolarGridHelper) ComputeFaceNormals() {
	g.Object.Call("computeFaceNormals")
}

func (g PolarGridHelper) ComputeVertexNormals(areaWeighted bool) {
	g.Object.Call("computeVertexNormals", areaWeighted)
}

func (g PolarGridHelper) ComputeFlatVertexNormals() {
	g.Object.Call("computeFlatVertexNormals")
}

func (g PolarGridHelper) ComputeMorphNormals() {
	g.Object.Call("computeMorphNormals")
}

func (g PolarGridHelper) ComputeLineDistances() {
	g.Object.Call("computeLineDistances")
}

func (g PolarGridHelper) ComputeBoundingBox() {
	g.Object.Call("computeBoundingBox")
}

func (g PolarGridHelper) ComputeBoundingSphere() {
	g.Object.Call("computeBoundingSphere")
}

func (g PolarGridHelper) Merge(geometry Geometry, matrix Matrix4, materialIndexOffset float64) {
	g.Object.Call("merge", geometry, matrix, materialIndexOffset)
}

func (g PolarGridHelper) MergeMesh(mesh Mesh) {
	g.Object.Call("mergeMesh", mesh.getInternalObject())
}

func (g PolarGridHelper) MergeVertices() {
	g.Object.Call("mergeVertices")
}

func (g PolarGridHelper) SortFacesByMaterialIndex() {
	g.Object.Call("sortFacesByMaterialIndex")
}

func (g PolarGridHelper) ToJSON() interface{} {
	return g.Object.Call("toJSON")
}

// func (g PolarGridHelper) Clone() PolarGridHelper {
// 	return g.Object.Call("clone")
// }

func (g PolarGridHelper) Copy(source Object3D, recursive bool) *PolarGridHelper {
	return &PolarGridHelper{Object: g.getInternalObject().Call("copy", source.getInternalObject(), recursive)}
}

func (g PolarGridHelper) Dispose() {
	g.Object.Call("dispose")
}

func (g PolarGridHelper) getInternalObject() *js.Object {
	return g.Object
}
