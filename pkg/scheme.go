package scheme

// #cgo CFLAGS: -I/opt/via/include/guile/2.2
// #cgo LDFLAGS: -L/opt/via/lib -lguile-2.2 -lgmp -lunistring -lffi -lm -lltdl -ldl -lcrypt -lgc
// #include "scheme.h"
// static void init() {
// scm_c_define_gsubr (s_scm_via_build, 0, 0, 0, (scm_t_subr) scm_via_build);;
// scm_c_export("build", NULL);
// }
import "C"
import (
	"fmt"
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

func (s SCM) Bool() bool {
	if C.scm_to_bool(s.box) == 1 {
		return true
	}
	return false
}

// IsBool returns true if SCM is a boolean
func (s SCM) IsBool() bool {
	if C.scm_is_bool(s.box) == 1 {
		return true
	}
	return false
}

// IsString returns true if SCM is a string
func (s SCM) IsString() bool {
	if C.scm_is_string(s.box) == 1 {
		return true
	}
	return false
}

// String provides stringer interface
func (s SCM) ToString() string {
	if !s.IsString() {
		return ""
	}
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

// AddToLoadPath add's path to %load-path
func AddToLoadPath(path string) SCM {
	scm := fmt.Sprintf(`(add-to-load-path "%s")`, path)
	fmt.Println(scm)
	return Eval(scm)
}

// Repl starts a new guile REPL
// FIXME: don't hardcode socket path
func Repl() SCM {
	Eval("(use-modules (system repl server))")
	return Eval(`(run-server
	(make-unix-domain-server-socket #:path "/tmp/go-scheme.socket"))`)
}

// Enter starts a console REPL server
func Enter() {
	C.scm_shell(0, nil)
}
