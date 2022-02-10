package gthree

import (
	"fmt"

	"github.com/bep/gr/support"
	"github.com/gopherjs/gopherjs/js"
)

var (
	three = js.Global.Get("THREE")
)

func init() {
	if three == js.Undefined {
		// Require as a fallback
		if _, err := support.Require("three"); err != nil {
			panic(fmt.Sprintf("Couldn't find three.js"))
		}
	}
}

// getModule gets a module or property added to THREE or globalThis.
func getModule(namespace string) *js.Object {
	mod := three.Get(namespace)
	if mod != js.Undefined {
		return mod
	}
	mod = js.Global.Get(namespace)
	if mod != js.Undefined {
		return mod
	}
	panic("three:failed to get " + namespace + " namespace from THREE and global namespace")
}
