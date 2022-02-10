package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/soypat/gthree"
)

var (
	scene    *gthree.Scene
	camera   gthree.PerspectiveCamera
	renderer gthree.WebGLRenderer
	mesh     *gthree.Mesh
)

func main() {
	document := js.Global.Get("document")
	windowWidth := js.Global.Get("innerWidth").Float()
	windowHeight := js.Global.Get("innerHeight").Float()
	devicePixelRatio := js.Global.Get("devicePixelRatio").Float()

	camera = gthree.NewPerspectiveCamera(70, windowWidth/windowHeight, 1, 1000)
	camera.Position.Set(0, 0, 400)

	scene = gthree.NewScene()

	light := gthree.NewDirectionalLight(gthree.NewColor("white"), 1)
	light.Position.Set(256, 256, 256).Normalize()
	scene.Add(light)

	ambLight := gthree.NewAmbientLight(gthree.NewColorHex(0xbbbbbb), 0.4)
	scene.Add(ambLight)

	renderer = gthree.NewWebGLRenderer()
	renderer.SetPixelRatio(devicePixelRatio)
	renderer.SetSize(windowWidth, windowHeight, true)
	document.Get("body").Call("appendChild", renderer.Get("domElement"))

	// Create cube
	geometry := gthree.NewBoxGeometry(&gthree.BoxGeometryParameters{
		Width:  128,
		Height: 128,
		Depth:  128,
	})

	// geometry2 := three.NewCircleGeometry(three.CircleGeometryParameters{
	// 	Radius:      50,
	// 	Segments:    20,
	// 	ThetaStart:  0,
	// 	ThetaLength: 2,
	// })

	materialParams := gthree.NewMaterialParameters()
	materialParams.Color = gthree.NewColor("blue")
	// materialParams.FlatShading = false
	materialParams.Side = gthree.FrontSide
	//material := three.NewMeshBasicMaterial(materialParams)
	material := gthree.NewMeshLambertMaterial(materialParams)
	//material := three.NewMeshPhongMaterial(materialParams)
	mesh = gthree.NewMesh(geometry, material)

	scene.Add(mesh)

	animate()
}

func animate() {
	pos := mesh.Object.Get("rotation")
	pos.Set("x", pos.Get("x").Float()+float64(0.01))
	pos.Set("y", pos.Get("y").Float()+float64(0.01))

	renderer.Render(scene, camera)

	// Best practice (soypat's opinion) to request frame after work is done.
	js.Global.Call("requestAnimationFrame", animate)
}
