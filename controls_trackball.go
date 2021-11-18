package three

import "github.com/gopherjs/gopherjs/js"

type TrackballControls struct {
	*js.Object

	DynamicDampingFactor float64 `js:"dynamicDampingFactor"`
	Enabled              bool    `js:"enabled"`

	// How far you can zoom out. Default is Infinity.
	MaxDistance float64 `js:"maxDistance"`
	// How far you can zoom in. Default is 0.
	MinDistance float64 `js:"minDistance"`
	// This object contains references to the mouse actions used by the controls.
	// * LEFT is assigned with THREE.MOUSE.ROTATE
	// * MIDDLE is assigned with THREE.MOUSE.ZOOM
	// * RIGHT is assigned with THREE.MOUSE.PAN
	MouseButtons *js.Object `js:"mouseButtons"`
	NoPan        bool       `js:"noPan"`
	NoRotate     bool       `js:"noRotate"`
	NoZoom       bool       `js:"noZoom"`
	Camera       Camera     `js:"object"`
	PanSpeed     float64    `js:"panSpeed"`
	RotateSpeed  float64    `js:"rotateSpeed"`
	// Whether or not damping is disabled. Default is false.
	StaticMoving bool    `js:"staticMoving"`
	ZoomSpeed    float64 `js:"zoomSpeed"`
}

// NewTrackballControls instances a TrackballControls. Requires TrackballControls to be
// added at a global level. See example https://github.com/mrdoob/three.js/blob/master/examples/jsm/controls/TrackballControls.js.
func NewTrackballControls(camera Camera, domElement *js.Object) TrackballControls {
	trackball := js.Global.Get("TrackballControls")
	if trackball == js.Undefined {
		panic("TrackballControls undefined at global level. Please add it somehow, god help you.")
	}
	if domElement == nil || domElement == js.Undefined {
		panic("domElement must be defined")
	}
	return TrackballControls{
		Object: trackball.New(camera, domElement),
	}
}

// Ensures the controls stay in the range [minDistance, maxDistance]. Called by update().
func (t TrackballControls) CheckDistances() {
	t.Object.Call("checkDistances")
}

// Should be called if the controls is no longer required.
func (t TrackballControls) Dispose() {
	t.Object.Call("dispose")
}

// Should be called if the application window is resized.
func (t TrackballControls) HandleResizes() {
	t.Object.Call("handleResizes")
}

// Performs panning if necessary. Called by update().
func (t TrackballControls) PanCamera() {
	t.Object.Call("panCamera")
}

// Resets the controls to its initial state.
func (t TrackballControls) Reset() {
	t.Object.Call("reset")
}

// Rotates the camera if necessary. Called by update().
func (t TrackballControls) RotateCamera() {
	t.Object.Call("rotateCamera")
}

// Updates the controls. Usually called in the animation loop.
func (t TrackballControls) Update() {
	t.Object.Call("update")
}

// Performs zooming if necessary. Called by update().
func (t TrackballControls) ZoomCamera() {
	t.Object.Call("zoomCamera")
}
