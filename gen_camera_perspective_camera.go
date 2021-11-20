package three

// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// camera_method_generator -typeName PerspectiveCamera -typeSlug perspective_camera

import "github.com/gopherjs/gopherjs/js"

var _ Camera = PerspectiveCamera{}

// Return a new camera with the same properties as this one.
func (obj PerspectiveCamera) Clone() PerspectiveCamera {
	return PerspectiveCamera{
		Object: obj.Call("clone"),
	}
}

// Copy the properties from the source camera into this one.
func (obj PerspectiveCamera) Copy(src PerspectiveCamera, recursive bool) PerspectiveCamera {
	return PerspectiveCamera{
		Object: obj.Call("copy", src, recursive),
	}
}

// Returns a Vector3 representing the world space direction 
// in which the camera is looking. (Note: A camera looks down its local, negative z-axis).
func (obj PerspectiveCamera) GetWorldDirection(target Vector3) Vector3 {
	return Vector3{
		Object: obj.Call("getWorldDirection", target),
	}
}

func (obj PerspectiveCamera) getInternalObject() *js.Object {
	return obj.Object
}
