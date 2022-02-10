## Archived
This package has been archived. I now favor developing with syscall/js package over gopherjs due to the near complete support it has for `unsafe` and other features of the Go language. If you are interested in a maintained, WASM-ready version please look at https://github.com/soypat/three

# three

GopherJS bindings for [three.js](https://threejs.org/). **Still a WIP.**

Keep in mind parts of these bindings are 4 years old and were not programmed by me :). I'm working hard-ish
to get everything back to speed with the latest version of three.js.

## Examples

* [examples directory](./examples) contains basic examples. There is a [earth with trackball controls](./examples/earth) example for camera manipulation.
* Example repo: https://github.com/soypat/threejs-golang-example of what package structure could look like for an gopherjs project.
* https://github.com/soypat/tiny-ahrsim: in the loop simulation of IMU attitude estimation on a Raspberry Pi Pico.
