package three

import "github.com/gopherjs/gopherjs/js"

// Vector3 - represents a Vector3.
type Vector3 struct {
	*js.Object
	X float64 `js:"x"`
	Y float64 `js:"y"`
	Z float64 `js:"z"`
}

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{
		Object: three.Get("Vector3").New(x, y, z),
	}
}

func (v Vector3) Set(x, y, z float64) Vector3 {
	v.Call("set", x, y, z)
	return v
}

func (v Vector3) SetX(x float64) Vector3 {
	v.Call("setX", x)
	return v
}

func (v Vector3) SetY(y float64) Vector3 {
	v.Call("setY", y)
	return v
}
func (v Vector3) SetZ(z float64) Vector3 {
	v.Call("setZ", z)
	return v
}

func (v Vector3) Coords() (x, y, z float64) {
	x = v.Object.Get("x").Float()
	y = v.Object.Get("y").Float()
	z = v.Object.Get("z").Float()
	return
}

func (v Vector3) Normalize() Vector3 {
	v.Call("normalize")
	return v
}
