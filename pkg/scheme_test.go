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
	got, err := Eval("%load-path")
	Tests{
		{
			Expect: nil,
			Got:    err,
		},
		{
			Expect: true,
			Got:    got.IsList(),
		},
	}.Equals(t)
}

func TestLoadPath(t *testing.T) {
	got, err := Eval("socket-file")
	Tests{
		{
			Expect: nil,
			Got:    err,
		},
		{
			Expect: "/tmp/go-scheme.socket",
			Got:    got.ToString(),
		},
	}.Equals(t)
}

func TestEval(t *testing.T) {
	got, err := Eval("(version)")
	Tests{
		{
			Expect: nil,
			Got:    err,
		},
		{
			Expect: "2.2.4",
			Got:    got.ToString(),
		},
	}.Equals(t)

}
