package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// object3d_method_generator -typeName AmbientLight -typeSlug ambient_light

import "github.com/gopherjs/gopherjs/js"

func (obj *AmbientLight) ApplyMatrix(matrix *Matrix4) {
	obj.Call("applyMatrix", matrix)
}

func (obj *AmbientLight) Add(m Object3D) {
	obj.Object.Call("add", m)
}

func (obj *AmbientLight) Remove(m *js.Object) {
	obj.Object.Call("remove", m)
}

func (obj *AmbientLight) GetObjectById(id int) *js.Object {
	return obj.Call("getObjectById", id)
}

// func (obj *AmbientLight) Copy() *AmbientLight {
// 	return &AmbientLight{Object: obj.getInternalObject().Call("copy")}
// }

func (obj *AmbientLight) ToJSON() interface{} {
	return obj.Object.Call("toJSON").Interface()
}

func (obj *AmbientLight) getInternalObject() *js.Object {
	return obj.Object
}

func (obj *AmbientLight) UpdateMatrix() {
	obj.Call("updateMatrix")
}
