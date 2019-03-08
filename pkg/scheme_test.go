package scheme

import (
	. "github.com/mrosset/test"
	"testing"
)

func TestVersion(t *testing.T) {
	Test{
		Expect: "2.2.4",
		Got:    Version().ToString(),
	}.Equals(t)
}

func TestList(t *testing.T) {
	Test{
		Expect: true,
		Got:    Eval("%load-path").IsList(),
	}.Equals(t)
}

func TestLoadPath(t *testing.T) {
	Test{
		Expect: "/tmp/go-scheme.socket",
		Got:    Eval("socket-file").ToString(),
	}.Equals(t)
}
