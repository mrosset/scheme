package main

// #include <libguile.h>
// #cgo pkg-config: guile-2.2
import "C"
import (
	"github.com/mrosset/scheme/pkg"
)

func main() {
	C.scm_init_guile()
	// scheme.Eval("(build)")
	scheme.Enter()
}
