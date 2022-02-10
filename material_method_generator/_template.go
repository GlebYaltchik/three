package gthree

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// material_method_generator -materialName {{ .Type }} -materialSlug {{ .Slug }}

import (
	"github.com/fatih/structs"
	"github.com/gopherjs/gopherjs/js"
)

var _ Material = {{ .Type }}{}

func (m {{ .Type }}) OnBeforeCompile() {
	m.Call("onBeforeCompile")
}

func (m {{ .Type }}) SetValues(values MaterialParameters) {
	m.Call("setValues", structs.Map(values))
}

func (m {{ .Type }}) ToJSON(meta interface{}) interface{} {
	return m.Call("toJSON", meta)
}

func (m {{ .Type }}) Clone() {
	m.Call("clone")
}

func (m {{ .Type }}) Copy(source Object3D) {
	m.Call("copy", source)
}

func (m {{ .Type }}) Dispose() {
	m.Call("dispose")
}

func (m {{ .Type }}) getInternalObject() *js.Object {
	return m.Object
}
