package scheme

import (
	"testing"
)

func TestVersion(t *testing.T) {
	var (
		expect = "3.0.7"
		got    = Version().ToString()
	)
	if expect != got {
		t.Errorf("Expect %s got %s", expect, got)
	}
}

func TestList(t *testing.T) {
	var (
		got, err = EvalString("%load-path")
	)
	if err != nil {
		t.Fatal(err)
	}
	if !got.IsList() {
		t.Errorf("Expext true got %t", got.IsList())
	}
}

// func TestSocketPath(t *testing.T) {
//	got, err := Eval("socket-file")
//	if err != nil {
//		t.Fatal(err)
//	}
//	Test{
//		Expect: "/tmp/go-scheme.socket",
//		Got:    got.ToString(),
//	}.Equals(t)
// }

func TestEval(t *testing.T) {
	var (
		got, err = EvalString("(version)")
		expect   = "3.0.7"
	)
	if err != nil {
		t.Fatal(err)
	}
	if got.ToString() != expect {
		t.Errorf("Expect %s got %s", got.ToString(), expect)
	}
}
