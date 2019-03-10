package scheme

// #cgo pkg-config: guile-2.2
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
	"os"
	"path/filepath"
	"unsafe"
)

var (
	LoadPath = filepath.Join(os.Getenv("GOPATH"), "src/github.com/mrosset/scheme/scm")
)

func init() {
	C.scm_init_guile()
	C.init()
	AddToLoadPath(LoadPath)
	UseModule("go server")
	UseModule("go eval")
}

// SCM provides a guile SCM type
type SCM struct {
	box C.SCM
}

// NewSCM returns a new initialized SCM type
func newSCM(scm C.SCM) SCM {
	return SCM{scm}
}

// Eval string returning a SCM
func Eval(exp string) (SCM, error) {
	var (
		ce   = C.CString(exp)
		cm   = C.CString("go eval")
		cp   = C.CString("go-eval")
		proc = C.scm_c_public_ref(cm, cp)
	)
	defer C.free(unsafe.Pointer(ce))
	defer C.free(unsafe.Pointer(cm))
	defer C.free(unsafe.Pointer(cp))
	res := C.scm_call_1(proc, C.scm_from_locale_string(ce))
	// n := C.scm_to_int(1)
	arg0 := C.scm_list_ref(res, C.scm_from_int(0))
	arg1 := C.scm_list_ref(res, C.scm_from_int(1))
	if C.scm_is_string(arg1) == 1 {
		return newSCM(arg0), fmt.Errorf("%s", newSCM(arg1).ToString())
	}
	return newSCM(arg0), nil
}

// Version returns guile scheme version
func Version() SCM {
	v, _ := Eval("(version)")
	return v
}

func evalstring(exp string) SCM {
	ce := C.CString(exp)
	defer C.free(unsafe.Pointer(ce))
	return newSCM(C.scm_c_eval_string(ce))
}

// AddToLoadPath add's path to %load-path
func AddToLoadPath(path string) SCM {
	exp := fmt.Sprintf(`(add-to-load-path "%s")`, path)
	return evalstring(exp)
}

// UseModule loads guile module
func UseModule(module string) {
	cs := C.CString(module)
	C.scm_c_use_module(cs)
	C.free(unsafe.Pointer(cs))
}

// Repl starts a new guile REPL
func Repl() (SCM, error) {
	return Eval("(server-start)")
}

// Enter starts a console REPL server
func Enter() {
	C.scm_shell(0, nil)
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

func (s SCM) IsList() bool {
	return newSCM(C.scm_list_p(s.box)).Bool()
}

func (s SCM) ToSlice() string {
	return ""
}
