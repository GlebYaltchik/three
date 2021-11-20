package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// object3d_method_generator -typeName Sprite -typeSlug sprite

import "github.com/gopherjs/gopherjs/js"

// Compile-time check that this type implements Object3D interface.
var _ Object3D = &Sprite{}

func (obj *Sprite) ApplyMatrix4(matrix *Matrix4) {
	obj.Call("applyMatrix4", matrix)
}

func (obj *Sprite) Add(m Object3D) {
	obj.Object.Call("add", m)
}

func (obj *Sprite) Remove(m *js.Object) {
	obj.Object.Call("remove", m)
}

func (obj *Sprite) GetObjectById(id int) *js.Object {
	return obj.Call("getObjectById", id)
}

// func (obj *Sprite) Copy() *Sprite {
// 	return &Sprite{Object: obj.getInternalObject().Call("copy")}
// }

func (obj *Sprite) ToJSON() interface{} {
	return obj.Object.Call("toJSON").Interface()
}

func (obj *Sprite) getInternalObject() *js.Object {
	return obj.Object
}

func (obj *Sprite) UpdateMatrix() {
	obj.Call("updateMatrix")
}
