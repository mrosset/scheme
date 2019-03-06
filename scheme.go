package scheme

// #cgo pkg-config: guile-2.2
// #include "scheme.h"
// static void init() {
// scm_c_define_gsubr (s_scm_via_build, 0, 0, 0, (scm_t_subr) scm_via_build);;
// scm_c_export("build", NULL);
// }
import "C"
import (
	"unsafe"
)

func init() {
	C.scm_init_guile()
	C.init()
}

// SCM provides a guile SCM type
type SCM struct {
	box C.SCM
}

// NewSCM returns a new initialized SCM type
func newSCM(scm C.SCM) SCM {
	return SCM{scm}
}

func (s SCM) String() string {
	cs := C.scm_to_locale_string(s.box)
	defer C.free(unsafe.Pointer(cs))
	return C.GoString(cs)
}

// Eval string returning a SCM
func Eval(expr string) SCM {
	var (
		cs  = C.CString(expr)
		res = C.scm_c_eval_string(cs)
	)
	defer C.free(unsafe.Pointer(cs))
	return newSCM(res)
}

// Version returns guile scheme version
func Version() SCM {
	return Eval("(version)")
}

// Repl starts a new guile REPL
func Repl() SCM {
	Eval("(use-modules (system repl server))")
	return Eval(`(run-server
	(make-unix-domain-server-socket #:path "/tmp/go-scheme.socket"))`)
}

// Enter starts a console REPL server
func Enter() {
	C.scm_shell(0, nil)
}
